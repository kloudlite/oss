// DO NOT EDIT. generated by "github.com/kloudlite/api/cmd/struct-json-path"

package field_constants

// constant vars generated for struct App
const (
	AppCiBuildId                      = "ciBuildId"
	AppEnabled                        = "enabled"
	AppSpec                           = "spec"
	AppSpecContainers                 = "spec.containers"
	AppSpecDisplayName                = "spec.displayName"
	AppSpecFreeze                     = "spec.freeze"
	AppSpecHpa                        = "spec.hpa"
	AppSpecHpaEnabled                 = "spec.hpa.enabled"
	AppSpecHpaMaxReplicas             = "spec.hpa.maxReplicas"
	AppSpecHpaMinReplicas             = "spec.hpa.minReplicas"
	AppSpecHpaThresholdCpu            = "spec.hpa.thresholdCpu"
	AppSpecHpaThresholdMemory         = "spec.hpa.thresholdMemory"
	AppSpecIntercept                  = "spec.intercept"
	AppSpecInterceptDeviceHostSuffix  = "spec.intercept.deviceHostSuffix"
	AppSpecInterceptEnabled           = "spec.intercept.enabled"
	AppSpecInterceptPortMappings      = "spec.intercept.portMappings"
	AppSpecInterceptToDevice          = "spec.intercept.toDevice"
	AppSpecNodeSelector               = "spec.nodeSelector"
	AppSpecRegion                     = "spec.region"
	AppSpecReplicas                   = "spec.replicas"
	AppSpecRouter                     = "spec.router"
	AppSpecRouterBackendProtocol      = "spec.router.backendProtocol"
	AppSpecRouterBasicAuth            = "spec.router.basicAuth"
	AppSpecRouterBasicAuthEnabled     = "spec.router.basicAuth.enabled"
	AppSpecRouterBasicAuthSecretName  = "spec.router.basicAuth.secretName"
	AppSpecRouterBasicAuthUsername    = "spec.router.basicAuth.username"
	AppSpecRouterCors                 = "spec.router.cors"
	AppSpecRouterCorsAllowCredentials = "spec.router.cors.allowCredentials"
	AppSpecRouterCorsEnabled          = "spec.router.cors.enabled"
	AppSpecRouterCorsOrigins          = "spec.router.cors.origins"
	AppSpecRouterDomains              = "spec.router.domains"
	AppSpecRouterHttps                = "spec.router.https"
	AppSpecRouterHttpsClusterIssuer   = "spec.router.https.clusterIssuer"
	AppSpecRouterHttpsEnabled         = "spec.router.https.enabled"
	AppSpecRouterHttpsForceRedirect   = "spec.router.https.forceRedirect"
	AppSpecRouterIngressClass         = "spec.router.ingressClass"
	AppSpecRouterMaxBodySizeInMB      = "spec.router.maxBodySizeInMB"
	AppSpecRouterRateLimit            = "spec.router.rateLimit"
	AppSpecRouterRateLimitConnections = "spec.router.rateLimit.connections"
	AppSpecRouterRateLimitEnabled     = "spec.router.rateLimit.enabled"
	AppSpecRouterRateLimitRpm         = "spec.router.rateLimit.rpm"
	AppSpecRouterRateLimitRps         = "spec.router.rateLimit.rps"
	AppSpecRouterRoutes               = "spec.router.routes"
	AppSpecServiceAccount             = "spec.serviceAccount"
	AppSpecServices                   = "spec.services"
	AppSpecTolerations                = "spec.tolerations"
	AppSpecTopologySpreadConstraints  = "spec.topologySpreadConstraints"
)

// constant vars generated for struct ClusterManagedService
const (
	ClusterManagedServiceIsArchived                                              = "isArchived"
	ClusterManagedServiceOutput                                                  = "output"
	ClusterManagedServiceOutputCredentialsRef                                    = "output.credentialsRef"
	ClusterManagedServiceOutputCredentialsRefName                                = "output.credentialsRef.name"
	ClusterManagedServiceSpec                                                    = "spec"
	ClusterManagedServiceSpecMsvcSpec                                            = "spec.msvcSpec"
	ClusterManagedServiceSpecMsvcSpecNodeSelector                                = "spec.msvcSpec.nodeSelector"
	ClusterManagedServiceSpecMsvcSpecServiceTemplate                             = "spec.msvcSpec.serviceTemplate"
	ClusterManagedServiceSpecMsvcSpecServiceTemplateApiVersion                   = "spec.msvcSpec.serviceTemplate.apiVersion"
	ClusterManagedServiceSpecMsvcSpecServiceTemplateKind                         = "spec.msvcSpec.serviceTemplate.kind"
	ClusterManagedServiceSpecMsvcSpecServiceTemplateSpec                         = "spec.msvcSpec.serviceTemplate.spec"
	ClusterManagedServiceSpecMsvcSpecSharedSecret                                = "spec.msvcSpec.sharedSecret"
	ClusterManagedServiceSpecMsvcSpecTolerations                                 = "spec.msvcSpec.tolerations"
	ClusterManagedServiceSpecSharedSecret                                        = "spec.sharedSecret"
	ClusterManagedServiceSpecTargetNamespace                                     = "spec.targetNamespace"
	ClusterManagedServiceSyncedOutputSecretRef                                   = "syncedOutputSecretRef"
	ClusterManagedServiceSyncedOutputSecretRefApiVersion                         = "syncedOutputSecretRef.apiVersion"
	ClusterManagedServiceSyncedOutputSecretRefData                               = "syncedOutputSecretRef.data"
	ClusterManagedServiceSyncedOutputSecretRefImmutable                          = "syncedOutputSecretRef.immutable"
	ClusterManagedServiceSyncedOutputSecretRefKind                               = "syncedOutputSecretRef.kind"
	ClusterManagedServiceSyncedOutputSecretRefMetadata                           = "syncedOutputSecretRef.metadata"
	ClusterManagedServiceSyncedOutputSecretRefMetadataAnnotations                = "syncedOutputSecretRef.metadata.annotations"
	ClusterManagedServiceSyncedOutputSecretRefMetadataCreationTimestamp          = "syncedOutputSecretRef.metadata.creationTimestamp"
	ClusterManagedServiceSyncedOutputSecretRefMetadataDeletionGracePeriodSeconds = "syncedOutputSecretRef.metadata.deletionGracePeriodSeconds"
	ClusterManagedServiceSyncedOutputSecretRefMetadataDeletionTimestamp          = "syncedOutputSecretRef.metadata.deletionTimestamp"
	ClusterManagedServiceSyncedOutputSecretRefMetadataFinalizers                 = "syncedOutputSecretRef.metadata.finalizers"
	ClusterManagedServiceSyncedOutputSecretRefMetadataGenerateName               = "syncedOutputSecretRef.metadata.generateName"
	ClusterManagedServiceSyncedOutputSecretRefMetadataGeneration                 = "syncedOutputSecretRef.metadata.generation"
	ClusterManagedServiceSyncedOutputSecretRefMetadataLabels                     = "syncedOutputSecretRef.metadata.labels"
	ClusterManagedServiceSyncedOutputSecretRefMetadataManagedFields              = "syncedOutputSecretRef.metadata.managedFields"
	ClusterManagedServiceSyncedOutputSecretRefMetadataName                       = "syncedOutputSecretRef.metadata.name"
	ClusterManagedServiceSyncedOutputSecretRefMetadataNamespace                  = "syncedOutputSecretRef.metadata.namespace"
	ClusterManagedServiceSyncedOutputSecretRefMetadataOwnerReferences            = "syncedOutputSecretRef.metadata.ownerReferences"
	ClusterManagedServiceSyncedOutputSecretRefMetadataResourceVersion            = "syncedOutputSecretRef.metadata.resourceVersion"
	ClusterManagedServiceSyncedOutputSecretRefMetadataSelfLink                   = "syncedOutputSecretRef.metadata.selfLink"
	ClusterManagedServiceSyncedOutputSecretRefMetadataUid                        = "syncedOutputSecretRef.metadata.uid"
	ClusterManagedServiceSyncedOutputSecretRefStringData                         = "syncedOutputSecretRef.stringData"
	ClusterManagedServiceSyncedOutputSecretRefType                               = "syncedOutputSecretRef.type"
)

// constant vars generated for struct Config
const (
	ConfigBinaryData = "binaryData"
	ConfigData       = "data"
	ConfigImmutable  = "immutable"
)

// constant vars generated for struct Environment
const (
	EnvironmentIsArchived                     = "isArchived"
	EnvironmentSpec                           = "spec"
	EnvironmentSpecRouting                    = "spec.routing"
	EnvironmentSpecRoutingMode                = "spec.routing.mode"
	EnvironmentSpecRoutingPrivateIngressClass = "spec.routing.privateIngressClass"
	EnvironmentSpecRoutingPublicIngressClass  = "spec.routing.publicIngressClass"
	EnvironmentSpecTargetNamespace            = "spec.targetNamespace"
)

// constant vars generated for struct ExternalApp
const (
	ExternalAppSpec                          = "spec"
	ExternalAppSpecIntercept                 = "spec.intercept"
	ExternalAppSpecInterceptDeviceHostSuffix = "spec.intercept.deviceHostSuffix"
	ExternalAppSpecInterceptEnabled          = "spec.intercept.enabled"
	ExternalAppSpecInterceptPortMappings     = "spec.intercept.portMappings"
	ExternalAppSpecInterceptToDevice         = "spec.intercept.toDevice"
	ExternalAppSpecRecord                    = "spec.record"
	ExternalAppSpecRecordType                = "spec.recordType"
)

// constant vars generated for struct ImagePullSecret
const (
	ImagePullSecretDockerConfigJson   = "dockerConfigJson"
	ImagePullSecretEnvironments       = "environments"
	ImagePullSecretFormat             = "format"
	ImagePullSecretGeneratedK8sSecret = "generatedK8sSecret"
	ImagePullSecretRegistryPassword   = "registryPassword"
	ImagePullSecretRegistryURL        = "registryURL"
	ImagePullSecretRegistryUsername   = "registryUsername"
)

// constant vars generated for struct ImportedManagedResource
const (
	ImportedManagedResourceManagedResourceRef          = "managedResourceRef"
	ImportedManagedResourceManagedResourceRefId        = "managedResourceRef.id"
	ImportedManagedResourceManagedResourceRefName      = "managedResourceRef.name"
	ImportedManagedResourceManagedResourceRefNamespace = "managedResourceRef.namespace"
	ImportedManagedResourceName                        = "name"
	ImportedManagedResourceSecretRef                   = "secretRef"
	ImportedManagedResourceSecretRefName               = "secretRef.name"
	ImportedManagedResourceSecretRefNamespace          = "secretRef.namespace"
)

// constant vars generated for struct ManagedResource
const (
	ManagedResourceEnabled                                                 = "enabled"
	ManagedResourceIsImported                                              = "isImported"
	ManagedResourceManagedServiceName                                      = "managedServiceName"
	ManagedResourceMresRef                                                 = "mresRef"
	ManagedResourceOutput                                                  = "output"
	ManagedResourceOutputCredentialsRef                                    = "output.credentialsRef"
	ManagedResourceOutputCredentialsRefName                                = "output.credentialsRef.name"
	ManagedResourceSpec                                                    = "spec"
	ManagedResourceSpecResourceNamePrefix                                  = "spec.resourceNamePrefix"
	ManagedResourceSpecResourceTemplate                                    = "spec.resourceTemplate"
	ManagedResourceSpecResourceTemplateApiVersion                          = "spec.resourceTemplate.apiVersion"
	ManagedResourceSpecResourceTemplateKind                                = "spec.resourceTemplate.kind"
	ManagedResourceSpecResourceTemplateMsvcRef                             = "spec.resourceTemplate.msvcRef"
	ManagedResourceSpecResourceTemplateMsvcRefApiVersion                   = "spec.resourceTemplate.msvcRef.apiVersion"
	ManagedResourceSpecResourceTemplateMsvcRefKind                         = "spec.resourceTemplate.msvcRef.kind"
	ManagedResourceSpecResourceTemplateMsvcRefName                         = "spec.resourceTemplate.msvcRef.name"
	ManagedResourceSpecResourceTemplateMsvcRefNamespace                    = "spec.resourceTemplate.msvcRef.namespace"
	ManagedResourceSpecResourceTemplateSpec                                = "spec.resourceTemplate.spec"
	ManagedResourceSyncedOutputSecretRef                                   = "syncedOutputSecretRef"
	ManagedResourceSyncedOutputSecretRefApiVersion                         = "syncedOutputSecretRef.apiVersion"
	ManagedResourceSyncedOutputSecretRefData                               = "syncedOutputSecretRef.data"
	ManagedResourceSyncedOutputSecretRefImmutable                          = "syncedOutputSecretRef.immutable"
	ManagedResourceSyncedOutputSecretRefKind                               = "syncedOutputSecretRef.kind"
	ManagedResourceSyncedOutputSecretRefMetadata                           = "syncedOutputSecretRef.metadata"
	ManagedResourceSyncedOutputSecretRefMetadataAnnotations                = "syncedOutputSecretRef.metadata.annotations"
	ManagedResourceSyncedOutputSecretRefMetadataCreationTimestamp          = "syncedOutputSecretRef.metadata.creationTimestamp"
	ManagedResourceSyncedOutputSecretRefMetadataDeletionGracePeriodSeconds = "syncedOutputSecretRef.metadata.deletionGracePeriodSeconds"
	ManagedResourceSyncedOutputSecretRefMetadataDeletionTimestamp          = "syncedOutputSecretRef.metadata.deletionTimestamp"
	ManagedResourceSyncedOutputSecretRefMetadataFinalizers                 = "syncedOutputSecretRef.metadata.finalizers"
	ManagedResourceSyncedOutputSecretRefMetadataGenerateName               = "syncedOutputSecretRef.metadata.generateName"
	ManagedResourceSyncedOutputSecretRefMetadataGeneration                 = "syncedOutputSecretRef.metadata.generation"
	ManagedResourceSyncedOutputSecretRefMetadataLabels                     = "syncedOutputSecretRef.metadata.labels"
	ManagedResourceSyncedOutputSecretRefMetadataManagedFields              = "syncedOutputSecretRef.metadata.managedFields"
	ManagedResourceSyncedOutputSecretRefMetadataName                       = "syncedOutputSecretRef.metadata.name"
	ManagedResourceSyncedOutputSecretRefMetadataNamespace                  = "syncedOutputSecretRef.metadata.namespace"
	ManagedResourceSyncedOutputSecretRefMetadataOwnerReferences            = "syncedOutputSecretRef.metadata.ownerReferences"
	ManagedResourceSyncedOutputSecretRefMetadataResourceVersion            = "syncedOutputSecretRef.metadata.resourceVersion"
	ManagedResourceSyncedOutputSecretRefMetadataSelfLink                   = "syncedOutputSecretRef.metadata.selfLink"
	ManagedResourceSyncedOutputSecretRefMetadataUid                        = "syncedOutputSecretRef.metadata.uid"
	ManagedResourceSyncedOutputSecretRefStringData                         = "syncedOutputSecretRef.stringData"
	ManagedResourceSyncedOutputSecretRefType                               = "syncedOutputSecretRef.type"
)

// constant vars generated for struct ManagedResourceRef
const (
	ManagedResourceRefName      = "name"
	ManagedResourceRefNamespace = "namespace"
)

// constant vars generated for struct ResourceMapping
const (
	ResourceMappingBaseEntity                  = "BaseEntity"
	ResourceMappingBaseEntityCreationTime      = "BaseEntity.creationTime"
	ResourceMappingBaseEntityId                = "BaseEntity.id"
	ResourceMappingBaseEntityMarkedForDeletion = "BaseEntity.markedForDeletion"
	ResourceMappingBaseEntityRecordVersion     = "BaseEntity.recordVersion"
	ResourceMappingBaseEntityUpdateTime        = "BaseEntity.updateTime"
	ResourceMappingResourceHeirarchy           = "resourceHeirarchy"
	ResourceMappingResourceName                = "resourceName"
	ResourceMappingResourceNamespace           = "resourceNamespace"
	ResourceMappingResourceType                = "resourceType"
)

// constant vars generated for struct Router
const (
	RouterEnabled                  = "enabled"
	RouterSpec                     = "spec"
	RouterSpecBackendProtocol      = "spec.backendProtocol"
	RouterSpecBasicAuth            = "spec.basicAuth"
	RouterSpecBasicAuthEnabled     = "spec.basicAuth.enabled"
	RouterSpecBasicAuthSecretName  = "spec.basicAuth.secretName"
	RouterSpecBasicAuthUsername    = "spec.basicAuth.username"
	RouterSpecCors                 = "spec.cors"
	RouterSpecCorsAllowCredentials = "spec.cors.allowCredentials"
	RouterSpecCorsEnabled          = "spec.cors.enabled"
	RouterSpecCorsOrigins          = "spec.cors.origins"
	RouterSpecDomains              = "spec.domains"
	RouterSpecHttps                = "spec.https"
	RouterSpecHttpsClusterIssuer   = "spec.https.clusterIssuer"
	RouterSpecHttpsEnabled         = "spec.https.enabled"
	RouterSpecHttpsForceRedirect   = "spec.https.forceRedirect"
	RouterSpecIngressClass         = "spec.ingressClass"
	RouterSpecMaxBodySizeInMB      = "spec.maxBodySizeInMB"
	RouterSpecRateLimit            = "spec.rateLimit"
	RouterSpecRateLimitConnections = "spec.rateLimit.connections"
	RouterSpecRateLimitEnabled     = "spec.rateLimit.enabled"
	RouterSpecRateLimitRpm         = "spec.rateLimit.rpm"
	RouterSpecRateLimitRps         = "spec.rateLimit.rps"
	RouterSpecRoutes               = "spec.routes"
)

// constant vars generated for struct Secret
const (
	SecretData            = "data"
	SecretFor             = "for"
	SecretForName         = "for.name"
	SecretForNamespace    = "for.namespace"
	SecretForRefId        = "for.refId"
	SecretForResourceType = "for.resourceType"
	SecretImmutable       = "immutable"
	SecretIsReadyOnly     = "isReadyOnly"
	SecretStringData      = "stringData"
	SecretType            = "type"
)

// constant vars generated for struct SecretCreatedFor
const (
	SecretCreatedForName         = "name"
	SecretCreatedForNamespace    = "namespace"
	SecretCreatedForRefId        = "refId"
	SecretCreatedForResourceType = "resourceType"
)

// constant vars generated for struct ServiceBinding
const (
	ServiceBindingSpec                    = "spec"
	ServiceBindingSpecGlobalIP            = "spec.globalIP"
	ServiceBindingSpecHostname            = "spec.hostname"
	ServiceBindingSpecPorts               = "spec.ports"
	ServiceBindingSpecServiceIP           = "spec.serviceIP"
	ServiceBindingSpecServiceRef          = "spec.serviceRef"
	ServiceBindingSpecServiceRefName      = "spec.serviceRef.name"
	ServiceBindingSpecServiceRefNamespace = "spec.serviceRef.namespace"
)

// constant vars generated for struct
const (
	AccountName                        = "accountName"
	ApiVersion                         = "apiVersion"
	ClusterName                        = "clusterName"
	CreatedBy                          = "createdBy"
	CreatedByUserEmail                 = "createdBy.userEmail"
	CreatedByUserId                    = "createdBy.userId"
	CreatedByUserName                  = "createdBy.userName"
	CreationTime                       = "creationTime"
	DisplayName                        = "displayName"
	EnvironmentName                    = "environmentName"
	Id                                 = "id"
	Kind                               = "kind"
	LastUpdatedBy                      = "lastUpdatedBy"
	LastUpdatedByUserEmail             = "lastUpdatedBy.userEmail"
	LastUpdatedByUserId                = "lastUpdatedBy.userId"
	LastUpdatedByUserName              = "lastUpdatedBy.userName"
	MarkedForDeletion                  = "markedForDeletion"
	Metadata                           = "metadata"
	MetadataAnnotations                = "metadata.annotations"
	MetadataCreationTimestamp          = "metadata.creationTimestamp"
	MetadataDeletionGracePeriodSeconds = "metadata.deletionGracePeriodSeconds"
	MetadataDeletionTimestamp          = "metadata.deletionTimestamp"
	MetadataFinalizers                 = "metadata.finalizers"
	MetadataGenerateName               = "metadata.generateName"
	MetadataGeneration                 = "metadata.generation"
	MetadataLabels                     = "metadata.labels"
	MetadataManagedFields              = "metadata.managedFields"
	MetadataName                       = "metadata.name"
	MetadataNamespace                  = "metadata.namespace"
	MetadataOwnerReferences            = "metadata.ownerReferences"
	MetadataResourceVersion            = "metadata.resourceVersion"
	MetadataSelfLink                   = "metadata.selfLink"
	MetadataUid                        = "metadata.uid"
	RecordVersion                      = "recordVersion"
	Status                             = "status"
	StatusCheckList                    = "status.checkList"
	StatusChecks                       = "status.checks"
	StatusIsReady                      = "status.isReady"
	StatusLastReadyGeneration          = "status.lastReadyGeneration"
	StatusLastReconcileTime            = "status.lastReconcileTime"
	StatusMessage                      = "status.message"
	StatusMessageItems                 = "status.message.items"
	StatusResources                    = "status.resources"
	SyncStatus                         = "syncStatus"
	SyncStatusAction                   = "syncStatus.action"
	SyncStatusError                    = "syncStatus.error"
	SyncStatusLastSyncedAt             = "syncStatus.lastSyncedAt"
	SyncStatusRecordVersion            = "syncStatus.recordVersion"
	SyncStatusState                    = "syncStatus.state"
	SyncStatusSyncScheduledAt          = "syncStatus.syncScheduledAt"
	UpdateTime                         = "updateTime"
)
