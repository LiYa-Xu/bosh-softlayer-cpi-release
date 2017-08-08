package instance

import (
	"fmt"
	"time"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"

	bsl "bosh-softlayer-cpi/softlayer/client"

	"bosh-softlayer-cpi/api"
	datatypes "github.com/softlayer/softlayer-go/datatypes"
	"strconv"
)

func (vg SoftlayerVirtualGuestService) getRootPassword(instance datatypes.Virtual_Guest) *string {
	passwords := (*instance.OperatingSystem).Passwords
	for _, password := range passwords {
		if *password.Username == rootUser {
			return password.Password
		}
	}

	return nil
}

func (vg SoftlayerVirtualGuestService) AttachEphemeralDisk(id int, diskSize int) error {
	return vg.softlayerClient.AttachSecondDiskToInstance(id, diskSize)
}

func (vg SoftlayerVirtualGuestService) AttachDisk(id int, diskID int) ([]byte, error) {
	ipAddress, found, err := vg.softlayerClient.GetNetworkStorageTarget(diskID, bsl.VOLUME_DETAIL_MASK)
	if err != nil {
		return []byte{}, bosherr.WrapErrorf(err, "Fetching disk target address with id '%d'", diskID)
	}

	if !found {
		return []byte{}, api.NewDiskNotFoundError(strconv.Itoa(diskID), false)
	}

	instance, found, err := vg.softlayerClient.GetInstance(id, bsl.INSTANCE_DETAIL_MASK)
	if err != nil {
		return []byte{}, bosherr.WrapErrorf(err, "Fetching instance details with id '%d'", id)
	}

	if !found {
		return []byte{}, api.NewVMNotFoundError(strconv.Itoa(id))
	}

	until := time.Now().Add(time.Duration(1) * time.Hour)
	_, err = vg.softlayerClient.AuthorizeHostToVolume(instance, diskID, until)
	if err != nil {
		return []byte{}, bosherr.WrapErrorf(err, "Authorizing vm with id '%d' to disk with id '%d'", id, diskID)
	}

	credential, found, err := vg.softlayerClient.GetAllowedHostCredential(id)
	if err != nil {
		return []byte{}, bosherr.WrapError(err, fmt.Sprintf("Getting iscsi host auth from virtual guest '%d'", id))
	}

	if !found {
		return []byte{}, api.NewHostHaveNotAllowedCredentialError(strconv.Itoa(id))
	}

	initiatorName := *credential.Name
	username := *credential.Credential.Username
	password := *credential.Credential.Password

	return []byte(fmt.Sprintf(`{"initiator_name":"%s","target":"%s","username":"%s","password":"%s" }`,
		initiatorName,
		ipAddress,
		username,
		password,
	)), nil
}

func (vg SoftlayerVirtualGuestService) AttachedDisks(id int) ([]string, error) {
	var attachedDisks []string
	attachedDisks, found, err := vg.softlayerClient.GetAllowedNetworkStorage(id)
	if err != nil {
		return attachedDisks, err
	}

	if !found {
		return attachedDisks, api.NewVMNotFoundError(strconv.Itoa(id))
	}

	return attachedDisks, nil
}

func (vg SoftlayerVirtualGuestService) DetachDisk(id int, diskID int) error {
	instance, found, err := vg.softlayerClient.GetInstance(id, bsl.INSTANCE_DETAIL_MASK)
	if err != nil {
		return bosherr.WrapErrorf(err, "Fetching instance details with id '%d'", id)
	}

	if !found {
		return api.NewVMNotFoundError(strconv.Itoa(id))
	}

	until := time.Now().Add(time.Duration(1) * time.Hour)
	_, err = vg.softlayerClient.DeauthorizeHostToVolume(instance, diskID, until)
	if err != nil {
		return bosherr.WrapErrorf(err, "De-Authorizing vm with id '%d' to disk with id '%d'", id, diskID)
	}

	return nil
}
