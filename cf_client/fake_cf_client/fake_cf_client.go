package fake_cf_client

import (
	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/api/apifakes"
	"code.cloudfoundry.org/cli/cf/api/organizations"
	"code.cloudfoundry.org/cli/cf/api/organizations/organizationsfakes"
	"code.cloudfoundry.org/cli/cf/api/quotas"
	"code.cloudfoundry.org/cli/cf/api/quotas/quotasfakes"
	"code.cloudfoundry.org/cli/cf/api/securitygroups"
	secgrouprun "code.cloudfoundry.org/cli/cf/api/securitygroups/defaults/running"
	secgrouprunfake "code.cloudfoundry.org/cli/cf/api/securitygroups/defaults/running/runningfakes"
	secgroupstag "code.cloudfoundry.org/cli/cf/api/securitygroups/defaults/staging"
	secgroupstagfake "code.cloudfoundry.org/cli/cf/api/securitygroups/defaults/staging/stagingfakes"
	"code.cloudfoundry.org/cli/cf/api/securitygroups/securitygroupsfakes"
	spacesbinder "code.cloudfoundry.org/cli/cf/api/securitygroups/spaces"
	spacesbinderfake "code.cloudfoundry.org/cli/cf/api/securitygroups/spaces/spacesfakes"
	"code.cloudfoundry.org/cli/cf/api/spacequotas"
	"code.cloudfoundry.org/cli/cf/api/spacequotas/spacequotasfakes"
	"code.cloudfoundry.org/cli/cf/api/spaces"
	"code.cloudfoundry.org/cli/cf/api/spaces/spacesfakes"
	"github.com/orange-cloudfoundry/terraform-provider-cloudfoundry/cf_client"
	"github.com/orange-cloudfoundry/terraform-provider-cloudfoundry/encryption"
	"github.com/orange-cloudfoundry/terraform-provider-cloudfoundry/encryption/fake_encryption"
)

type FakeCfClient struct {
	config                      cf_client.Config
	organizations               *organizationsfakes.FakeOrganizationRepository
	spaces                      *spacesfakes.FakeSpaceRepository
	securityGroups              *securitygroupsfakes.FakeSecurityGroupRepo
	serviceBrokers              *apifakes.FakeServiceBrokerRepository
	servicePlanVisibilities     *apifakes.FakeServicePlanVisibilityRepository
	spaceQuotas                 *spacequotasfakes.FakeSpaceQuotaRepository
	quotas                      *quotasfakes.FakeQuotaRepository
	buildpack                   *apifakes.FakeBuildpackRepository
	buildpackBits               *apifakes.FakeBuildpackBitsRepository
	securityGroupsSpaceBinder   *spacesbinderfake.FakeSecurityGroupSpaceBinder
	securityGroupsRunningBinder *secgrouprunfake.FakeSecurityGroupsRepo
	securityGroupsStagingBinder *secgroupstagfake.FakeSecurityGroupsRepo
	servicePlans                *apifakes.FakeServicePlanRepository
	decrypter                   encryption.Decrypter
	services                    *apifakes.FakeServiceRepository
}

func NewFakeCfClient() *FakeCfClient {
	client := &FakeCfClient{}
	client.Init()
	return client
}
func (c *FakeCfClient) GetClient() cf_client.Client {
	return c
}
func (c *FakeCfClient) Init() {
	c.config = cf_client.Config{
		ApiEndpoint: "http://fake.api.endpoint.com",
	}
	c.organizations = new(organizationsfakes.FakeOrganizationRepository)
	c.spaces = new(spacesfakes.FakeSpaceRepository)
	c.securityGroups = new(securitygroupsfakes.FakeSecurityGroupRepo)
	c.serviceBrokers = new(apifakes.FakeServiceBrokerRepository)
	c.servicePlanVisibilities = new(apifakes.FakeServicePlanVisibilityRepository)
	c.spaceQuotas = new(spacequotasfakes.FakeSpaceQuotaRepository)
	c.quotas = new(quotasfakes.FakeQuotaRepository)
	c.buildpack = new(apifakes.FakeBuildpackRepository)
	c.buildpackBits = new(apifakes.FakeBuildpackBitsRepository)
	c.securityGroupsSpaceBinder = new(spacesbinderfake.FakeSecurityGroupSpaceBinder)
	c.securityGroupsRunningBinder = new(secgrouprunfake.FakeSecurityGroupsRepo)
	c.securityGroupsStagingBinder = new(secgroupstagfake.FakeSecurityGroupsRepo)
	c.servicePlans = new(apifakes.FakeServicePlanRepository)
	c.services = new(apifakes.FakeServiceRepository)
	c.decrypter = fake_encryption.NewFakeDecrypter()
}
func (client FakeCfClient) Organizations() organizations.OrganizationRepository {
	return client.organizations
}

func (client FakeCfClient) Spaces() spaces.SpaceRepository {
	return client.spaces
}
func (client FakeCfClient) SecurityGroups() securitygroups.SecurityGroupRepo {
	return client.securityGroups
}
func (client FakeCfClient) ServiceBrokers() api.ServiceBrokerRepository {
	return client.serviceBrokers
}
func (client FakeCfClient) ServicePlanVisibilities() api.ServicePlanVisibilityRepository {
	return client.servicePlanVisibilities
}
func (client FakeCfClient) SpaceQuotas() spacequotas.SpaceQuotaRepository {
	return client.spaceQuotas
}

func (client FakeCfClient) Quotas() quotas.QuotaRepository {
	return client.quotas
}

func (client FakeCfClient) Config() cf_client.Config {
	return client.config
}
func (client FakeCfClient) Buildpack() api.BuildpackRepository {
	return client.buildpack
}
func (client FakeCfClient) BuildpackBits() api.BuildpackBitsRepository {
	return client.buildpackBits
}
func (client FakeCfClient) SecurityGroupsSpaceBinder() spacesbinder.SecurityGroupSpaceBinder {
	return client.securityGroupsSpaceBinder
}
func (client FakeCfClient) SecurityGroupsRunningBinder() secgrouprun.SecurityGroupsRepo {
	return client.securityGroupsRunningBinder
}
func (client FakeCfClient) SecurityGroupsStagingBinder() secgroupstag.SecurityGroupsRepo {
	return client.securityGroupsStagingBinder
}
func (client FakeCfClient) ServicePlans() api.ServicePlanRepository {
	return client.servicePlans
}
func (client FakeCfClient) Services() api.ServiceRepository {
	return client.services
}
func (client FakeCfClient) Decrypter() encryption.Decrypter {
	return client.decrypter
}

// get Fake call -------

func (client FakeCfClient) FakeOrganizations() *organizationsfakes.FakeOrganizationRepository {
	return client.organizations
}

func (client FakeCfClient) FakeSpaces() *spacesfakes.FakeSpaceRepository {
	return client.spaces
}
func (client FakeCfClient) FakeSecurityGroups() *securitygroupsfakes.FakeSecurityGroupRepo {
	return client.securityGroups
}
func (client FakeCfClient) FakeServiceBrokers() *apifakes.FakeServiceBrokerRepository {
	return client.serviceBrokers
}
func (client FakeCfClient) FakeServicePlanVisibilities() *apifakes.FakeServicePlanVisibilityRepository {
	return client.servicePlanVisibilities
}
func (client FakeCfClient) FakeSpaceQuotas() *spacequotasfakes.FakeSpaceQuotaRepository {
	return client.spaceQuotas
}

func (client FakeCfClient) FakeQuotas() *quotasfakes.FakeQuotaRepository {
	return client.quotas
}

func (client FakeCfClient) FakeBuildpack() *apifakes.FakeBuildpackRepository {
	return client.buildpack
}
func (client FakeCfClient) FakeBuildpackBits() *apifakes.FakeBuildpackBitsRepository {
	return client.buildpackBits
}
func (client FakeCfClient) FakeSecurityGroupsSpaceBinder() *spacesbinderfake.FakeSecurityGroupSpaceBinder {
	return client.securityGroupsSpaceBinder
}
func (client FakeCfClient) FakeSecurityGroupsRunningBinder() *secgrouprunfake.FakeSecurityGroupsRepo {
	return client.securityGroupsRunningBinder
}
func (client FakeCfClient) FakeSecurityGroupsStagingBinder() *secgroupstagfake.FakeSecurityGroupsRepo {
	return client.securityGroupsStagingBinder
}
func (client FakeCfClient) FakeServicePlans() *apifakes.FakeServicePlanRepository {
	return client.servicePlans
}
func (client FakeCfClient) FakeServices() *apifakes.FakeServiceRepository {
	return client.services
}
