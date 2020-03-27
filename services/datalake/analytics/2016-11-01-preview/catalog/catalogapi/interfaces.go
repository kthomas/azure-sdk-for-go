package catalogapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/datalake/analytics/2016-11-01-preview/catalog"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateCredential(ctx context.Context, accountName string, databaseName string, credentialName string, parameters catalog.DataLakeAnalyticsCatalogCredentialCreateParameters) (result autorest.Response, err error)
	CreateSecret(ctx context.Context, accountName string, databaseName string, secretName string, parameters catalog.DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters) (result autorest.Response, err error)
	DeleteAllSecrets(ctx context.Context, accountName string, databaseName string) (result autorest.Response, err error)
	DeleteCredential(ctx context.Context, accountName string, databaseName string, credentialName string, parameters *catalog.DataLakeAnalyticsCatalogCredentialDeleteParameters, cascade *bool) (result autorest.Response, err error)
	DeleteSecret(ctx context.Context, accountName string, databaseName string, secretName string) (result autorest.Response, err error)
	GetAssembly(ctx context.Context, accountName string, databaseName string, assemblyName string) (result catalog.USQLAssembly, err error)
	GetCredential(ctx context.Context, accountName string, databaseName string, credentialName string) (result catalog.USQLCredential, err error)
	GetDatabase(ctx context.Context, accountName string, databaseName string) (result catalog.USQLDatabase, err error)
	GetExternalDataSource(ctx context.Context, accountName string, databaseName string, externalDataSourceName string) (result catalog.USQLExternalDataSource, err error)
	GetPackage(ctx context.Context, accountName string, databaseName string, schemaName string, packageName string) (result catalog.USQLPackage, err error)
	GetProcedure(ctx context.Context, accountName string, databaseName string, schemaName string, procedureName string) (result catalog.USQLProcedure, err error)
	GetSchema(ctx context.Context, accountName string, databaseName string, schemaName string) (result catalog.USQLSchema, err error)
	GetSecret(ctx context.Context, accountName string, databaseName string, secretName string) (result catalog.USQLSecret, err error)
	GetTable(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string) (result catalog.USQLTable, err error)
	GetTablePartition(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, partitionName string) (result catalog.USQLTablePartition, err error)
	GetTableStatistic(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, statisticsName string) (result catalog.USQLTableStatistics, err error)
	GetTableType(ctx context.Context, accountName string, databaseName string, schemaName string, tableTypeName string) (result catalog.USQLTableType, err error)
	GetTableValuedFunction(ctx context.Context, accountName string, databaseName string, schemaName string, tableValuedFunctionName string) (result catalog.USQLTableValuedFunction, err error)
	GetView(ctx context.Context, accountName string, databaseName string, schemaName string, viewName string) (result catalog.USQLView, err error)
	GrantACL(ctx context.Context, accountName string, parameters catalog.ACLCreateOrUpdateParameters) (result autorest.Response, err error)
	GrantACLToDatabase(ctx context.Context, accountName string, databaseName string, parameters catalog.ACLCreateOrUpdateParameters) (result autorest.Response, err error)
	ListAcls(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.ACLListPage, err error)
	ListAclsComplete(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.ACLListIterator, err error)
	ListAclsByDatabase(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.ACLListPage, err error)
	ListAclsByDatabaseComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.ACLListIterator, err error)
	ListAssemblies(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLAssemblyListPage, err error)
	ListAssembliesComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLAssemblyListIterator, err error)
	ListCredentials(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLCredentialListPage, err error)
	ListCredentialsComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLCredentialListIterator, err error)
	ListDatabases(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLDatabaseListPage, err error)
	ListDatabasesComplete(ctx context.Context, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLDatabaseListIterator, err error)
	ListExternalDataSources(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLExternalDataSourceListPage, err error)
	ListExternalDataSourcesComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLExternalDataSourceListIterator, err error)
	ListPackages(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLPackageListPage, err error)
	ListPackagesComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLPackageListIterator, err error)
	ListProcedures(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLProcedureListPage, err error)
	ListProceduresComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLProcedureListIterator, err error)
	ListSchemas(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLSchemaListPage, err error)
	ListSchemasComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLSchemaListIterator, err error)
	ListTableFragments(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableFragmentListPage, err error)
	ListTableFragmentsComplete(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableFragmentListIterator, err error)
	ListTablePartitions(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTablePartitionListPage, err error)
	ListTablePartitionsComplete(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTablePartitionListIterator, err error)
	ListTables(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool, basic *bool) (result catalog.USQLTableListPage, err error)
	ListTablesComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool, basic *bool) (result catalog.USQLTableListIterator, err error)
	ListTablesByDatabase(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool, basic *bool) (result catalog.USQLTableListPage, err error)
	ListTablesByDatabaseComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool, basic *bool) (result catalog.USQLTableListIterator, err error)
	ListTableStatistics(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableStatisticsListPage, err error)
	ListTableStatisticsComplete(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableStatisticsListIterator, err error)
	ListTableStatisticsByDatabase(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableStatisticsListPage, err error)
	ListTableStatisticsByDatabaseComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableStatisticsListIterator, err error)
	ListTableStatisticsByDatabaseAndSchema(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableStatisticsListPage, err error)
	ListTableStatisticsByDatabaseAndSchemaComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableStatisticsListIterator, err error)
	ListTableTypes(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableTypeListPage, err error)
	ListTableTypesComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableTypeListIterator, err error)
	ListTableValuedFunctions(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableValuedFunctionListPage, err error)
	ListTableValuedFunctionsComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableValuedFunctionListIterator, err error)
	ListTableValuedFunctionsByDatabase(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableValuedFunctionListPage, err error)
	ListTableValuedFunctionsByDatabaseComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTableValuedFunctionListIterator, err error)
	ListTypes(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTypeListPage, err error)
	ListTypesComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLTypeListIterator, err error)
	ListViews(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLViewListPage, err error)
	ListViewsComplete(ctx context.Context, accountName string, databaseName string, schemaName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLViewListIterator, err error)
	ListViewsByDatabase(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLViewListPage, err error)
	ListViewsByDatabaseComplete(ctx context.Context, accountName string, databaseName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result catalog.USQLViewListIterator, err error)
	PreviewTable(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, maxRows *int64, maxColumns *int64) (result catalog.USQLTablePreview, err error)
	PreviewTablePartition(ctx context.Context, accountName string, databaseName string, schemaName string, tableName string, partitionName string, maxRows *int64, maxColumns *int64) (result catalog.USQLTablePreview, err error)
	RevokeACL(ctx context.Context, accountName string, parameters catalog.ACLDeleteParameters) (result autorest.Response, err error)
	RevokeACLFromDatabase(ctx context.Context, accountName string, databaseName string, parameters catalog.ACLDeleteParameters) (result autorest.Response, err error)
	UpdateCredential(ctx context.Context, accountName string, databaseName string, credentialName string, parameters catalog.DataLakeAnalyticsCatalogCredentialUpdateParameters) (result autorest.Response, err error)
	UpdateSecret(ctx context.Context, accountName string, databaseName string, secretName string, parameters catalog.DataLakeAnalyticsCatalogSecretCreateOrUpdateParameters) (result autorest.Response, err error)
}

var _ ClientAPI = (*catalog.Client)(nil)
