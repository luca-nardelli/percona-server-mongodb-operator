package stub

import (
	"testing"

	"github.com/Percona-Lab/percona-server-mongodb-operator/pkg/apis/cache/v1alpha1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestGetWiredTigerCacheSizeGB(t *testing.T) {
	memoryQuantity := resource.NewQuantity(gigaByte/2, resource.DecimalSI)
	assert.Equal(t, 0.25, getWiredTigerCacheSizeGB(memoryQuantity, 0.5))

	memoryQuantity = resource.NewQuantity(1*gigaByte, resource.DecimalSI)
	assert.Equal(t, 0.25, getWiredTigerCacheSizeGB(memoryQuantity, 0.5))

	memoryQuantity = resource.NewQuantity(4*gigaByte, resource.DecimalSI)
	assert.Equal(t, 1.5, getWiredTigerCacheSizeGB(memoryQuantity, 0.5))

	memoryQuantity = resource.NewQuantity(64*gigaByte, resource.DecimalSI)
	assert.Equal(t, 31.5, getWiredTigerCacheSizeGB(memoryQuantity, 0.5))

	memoryQuantity = resource.NewQuantity(128*gigaByte, resource.DecimalSI)
	assert.Equal(t, 63.5, getWiredTigerCacheSizeGB(memoryQuantity, 0.5))

	memoryQuantity = resource.NewQuantity(256*gigaByte, resource.DecimalSI)
	assert.Equal(t, 127.5, getWiredTigerCacheSizeGB(memoryQuantity, 0.5))
}

func TestAddPSMDBSpecDefaults(t *testing.T) {
	spec := v1alpha1.PerconaServerMongoDBSpec{}

	addPSMDBSpecDefaults(&spec)

	assert.Equal(t, defaultSize, spec.Size)
	assert.Equal(t, defaultImage, spec.Image)
	assert.NotNil(t, spec.MongoDB)
	assert.Equal(t, defaultStorageEngine, spec.MongoDB.StorageEngine)
	assert.NotNil(t, spec.MongoDB.WiredTiger)
	assert.Equal(t, defaultWiredTigerCacheSizeRatio, spec.MongoDB.WiredTiger.CacheSizeRatio)
}