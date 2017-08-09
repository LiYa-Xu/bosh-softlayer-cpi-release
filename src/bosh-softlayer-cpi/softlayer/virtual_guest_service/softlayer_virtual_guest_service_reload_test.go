package instance_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	fakeslclient "bosh-softlayer-cpi/softlayer/client/fakes"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	fakeuuid "github.com/cloudfoundry/bosh-utils/uuid/fakes"

	. "bosh-softlayer-cpi/softlayer/virtual_guest_service"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/sl"
)

var _ = Describe("Virtual Guest Service", func() {
	var (
		cli                 *fakeslclient.FakeClient
		uuidGen             *fakeuuid.FakeGenerator
		logger              boshlog.Logger
		virtualGuestService SoftlayerVirtualGuestService
	)

	BeforeEach(func() {
		cli = &fakeslclient.FakeClient{}
		uuidGen = &fakeuuid.FakeGenerator{}
		logger = boshlog.NewLogger(boshlog.LevelDebug)
		virtualGuestService = NewSoftLayerVirtualGuestService(cli, uuidGen, logger)
	})

	Describe("Call ReloadOS", func() {
		var (
			vmID         int
			stemcellID   int
			sshKeys      []int
			vmNamePrefix string
			domain       string
		)

		BeforeEach(func() {
			vmID = 12345678
			stemcellID = 22345678
			sshKeys = []int{342345, 42345}
			vmNamePrefix = "fake-prefix"
			domain = "unit-test"

			cli.ReloadInstanceReturns(
				nil,
			)
		})

		It("Reload instance successfully", func() {
			err := virtualGuestService.ReloadOS(vmID, stemcellID, sshKeys, vmNamePrefix, domain)
			Expect(err).NotTo(HaveOccurred())
			Expect(cli.ReloadInstanceCallCount()).To(Equal(1))
		})

		It("Return error if softLayerClient ReloadInstance call returns an error", func() {
			cli.ReloadInstanceReturns(
				errors.New("fake-client-error"),
			)

			err := virtualGuestService.ReloadOS(vmID, stemcellID, sshKeys, vmNamePrefix, domain)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-client-error"))
			Expect(cli.ReloadInstanceCallCount()).To(Equal(1))
		})
	})

	Describe("Call Edit", func() {
		var (
			vmID int

			updateVirtualGuest *datatypes.Virtual_Guest
		)

		BeforeEach(func() {
			vmID = 12345678
			updateVirtualGuest = &datatypes.Virtual_Guest{
				StartCpus: sl.Int(4),
				MaxMemory: sl.Int(1000),
			}

			cli.EditInstanceReturns(
				true,
				nil,
			)
		})

		It("Edit instance successfully", func() {
			err := virtualGuestService.Edit(vmID, updateVirtualGuest)
			Expect(err).NotTo(HaveOccurred())
			Expect(cli.EditInstanceCallCount()).To(Equal(1))
		})

		It("Return error if softLayerClient EditInstance call returns an error", func() {
			cli.EditInstanceReturns(
				true,
				errors.New("fake-client-error"),
			)

			err := virtualGuestService.Edit(vmID, updateVirtualGuest)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-client-error"))
			Expect(cli.EditInstanceCallCount()).To(Equal(1))
		})
	})
})
