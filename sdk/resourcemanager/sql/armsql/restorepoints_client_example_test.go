//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRestorePointsListByDatabase.json
func ExampleRestorePointsClient_NewListByDatabasePager_listDatabaseRestorePoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointsClient().NewListByDatabasePager("sqlcrudtest-6730", "sqlcrudtest-9007", "3481", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RestorePointListResult = armsql.RestorePointListResult{
		// 	Value: []*armsql.RestorePoint{
		// 		{
		// 			Name: to.Ptr("ContinuousRestorePoint"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6730/providers/Microsoft.Sql/servers/sqlcrudtest-9007/databases/3481/restorepoints/ContinuousRestorePoint"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armsql.RestorePointProperties{
		// 				EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-12T00:00:00Z"); return t}()),
		// 				RestorePointType: to.Ptr(armsql.RestorePointTypeCONTINUOUS),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DataWarehouseRestorePointsListByDatabase.json
func ExampleRestorePointsClient_NewListByDatabasePager_listDatawarehouseDatabaseRestorePoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointsClient().NewListByDatabasePager("Default-SQL-SouthEastAsia", "testserver", "testDatabase", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.RestorePointListResult = armsql.RestorePointListResult{
		// 	Value: []*armsql.RestorePoint{
		// 		{
		// 			Name: to.Ptr("131546477590000000"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131546477590000000"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.RestorePointProperties{
		// 				RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
		// 				RestorePointLabel: to.Ptr("mylabel1"),
		// 				RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("131553636140000000"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131553636140000000"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.RestorePointProperties{
		// 				RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T03:40:14Z"); return t}()),
		// 				RestorePointLabel: to.Ptr("mylabel2"),
		// 				RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("131553619750000000"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131553619750000000"),
		// 			Location: to.Ptr("japaneast"),
		// 			Properties: &armsql.RestorePointProperties{
		// 				RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T03:12:55Z"); return t}()),
		// 				RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRestorePointsPost.json
func ExampleRestorePointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRestorePointsClient().BeginCreate(ctx, "Default-SQL-SouthEastAsia", "testserver", "testDatabase", armsql.CreateDatabaseRestorePointDefinition{
		RestorePointLabel: to.Ptr("mylabel"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armsql.RestorePoint{
	// 	Name: to.Ptr("131546477590000000"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131546477590000000"),
	// 	Location: to.Ptr("japaneast"),
	// 	Properties: &armsql.RestorePointProperties{
	// 		RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
	// 		RestorePointLabel: to.Ptr("mylabel"),
	// 		RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRestorePointsGet.json
func ExampleRestorePointsClient_Get_getsADatabaseRestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testserver", "testDatabase", "131546477590000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armsql.RestorePoint{
	// 	Name: to.Ptr("ContinuousRestorePoint"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/ContinuousRestorePoint"),
	// 	Location: to.Ptr("japaneast"),
	// 	Properties: &armsql.RestorePointProperties{
	// 		EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
	// 		RestorePointType: to.Ptr(armsql.RestorePointTypeCONTINUOUS),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DataWarehouseRestorePointsGet.json
func ExampleRestorePointsClient_Get_getsADatawarehouseDatabaseRestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testserver", "testDatabase", "131546477590000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armsql.RestorePoint{
	// 	Name: to.Ptr("131546477590000000"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131546477590000000"),
	// 	Location: to.Ptr("japaneast"),
	// 	Properties: &armsql.RestorePointProperties{
	// 		RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00Z"); return t}()),
	// 		RestorePointLabel: to.Ptr("mylabel"),
	// 		RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRestorePointsDelete.json
func ExampleRestorePointsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRestorePointsClient().Delete(ctx, "Default-SQL-SouthEastAsia", "testserver", "testDatabase", "131546477590000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
