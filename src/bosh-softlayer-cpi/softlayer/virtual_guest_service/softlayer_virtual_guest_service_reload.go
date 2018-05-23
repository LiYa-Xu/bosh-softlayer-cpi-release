package instance

import (
	"bosh-softlayer-cpi/api"
	"github.com/softlayer/softlayer-go/datatypes"
	"strconv"
)

func (vg SoftlayerVirtualGuestService) ReloadOS(id int, stemcellID int, sshKeys []int, hostname string, domain string) error {
	vg.logger.Debug(softlayerVirtualGuestServiceLogTag, "Reloading instance '%d'.", id)
	return vg.softlayerClient.ReloadInstance(id, stemcellID, sshKeys, hostname, domain)
}

func (vg SoftlayerVirtualGuestService) Edit(id int, instance *datatypes.Virtual_Guest) error {
	vg.logger.Debug(softlayerVirtualGuestServiceLogTag, "Editing instance '%d'.", id)
	found, err := vg.softlayerClient.EditInstance(id, instance)
	if err != nil {
		return err
	}

	if !found {
		return api.NewVMNotFoundError(strconv.Itoa(id))
	}

	return nil
}

func (vg SoftlayerVirtualGuestService) UpdateInstanceUserData(id int, userData *string) error {
	vg.logger.Debug(softlayerVirtualGuestServiceLogTag, "Setting instance '%d' userData: %s", id, *userData)
	found, err := vg.softlayerClient.SetInstanceMetadata(id, userData)
	if err != nil {
		return err
	}

	if !found {
		return api.NewVMNotFoundError(strconv.Itoa(id))
	}

	return nil
}
