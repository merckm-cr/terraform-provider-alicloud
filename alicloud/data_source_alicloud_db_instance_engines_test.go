package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
)

func TestAccAlicloudDBEngines_base(t *testing.T) {
	rand := acctest.RandInt()
	ZoneIDConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"zone_id": `"${data.alicloud_zones.resources.zones.0.id}"`,
		}),
		fakeConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"zone_id": `"fake_zoneid"`,
		}),
	}

	ChargeTypeConf_Postpaid := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"instance_charge_type": `"PostPaid"`,
			"zone_id":              `"${data.alicloud_zones.resources.zones.0.id}"`,
		}),
	}
	ChargeTypeConf_Prepaid := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"instance_charge_type": `"PrePaid"`,
			"zone_id":              `"${data.alicloud_zones.resources.zones.0.id}"`,
		}),
	}
	EngineConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"engine": `"MySQL"`,
		}),
		fakeConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"engine": `"Fake"`,
		}),
	}
	EngineVersionConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"engine":         `"MySQL"`,
			"engine_version": `"5.6"`,
		}),
		fakeConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"engine":         `"MySQL"`,
			"engine_version": `"3.0"`,
		}),
	}
	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"instance_charge_type": `"PostPaid"`,
			"engine":               `"MySQL"`,
			"engine_version":       `"5.6"`,
			"zone_id":              `"${data.alicloud_zones.resources.zones.0.id}"`,
		}),
		fakeConfig: testAccCheckAlicloudDBEnginesDataSourceConfig(map[string]string{
			"zone_id":              `"${data.alicloud_zones.resources.zones.0.id}"`,
			"instance_charge_type": `"PostPaid"`,
			"engine":               `"MySQL"`,
			"engine_version":       `"3.0"`,
		}),
	}

	var existDBInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_engines.#":                CHECKSET,
			"instance_engines.0.engine":         CHECKSET,
			"instance_engines.0.zone_id":        CHECKSET,
			"instance_engines.0.engine_version": CHECKSET,
			"instance_engines.0.category":       CHECKSET,
		}
	}

	var fakeDBInstanceMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"instance_engines.#": "0",
		}
	}

	var DBInstanceCheckInfo = dataSourceAttr{
		resourceId:   "data.alicloud_db_instance_engines.resources",
		existMapFunc: existDBInstanceMapFunc,
		fakeMapFunc:  fakeDBInstanceMapFunc,
	}
	DBInstanceCheckInfo.dataSourceTestCheck(t, rand, ZoneIDConf, ChargeTypeConf_Postpaid, ChargeTypeConf_Prepaid, EngineConf, EngineVersionConf, allConf)
}

func testAccCheckAlicloudDBEnginesDataSourceConfig(attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	config := fmt.Sprintf(`
data "alicloud_zones" "resources" {
  available_resource_creation= "Rds"
}
data "alicloud_db_instance_engines" "resources" {
  %s
}
`, strings.Join(pairs, "\n  "))
	return config
}
