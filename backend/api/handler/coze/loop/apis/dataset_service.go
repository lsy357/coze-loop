// Copyright (c) 2025 coze-dev Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by hertz generator.

package apis

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/coze-dev/coze-loop/backend/kitex_gen/coze/loop/data/dataset/datasetservice"
)

var localDataSvc datasetservice.Client

// CreateDataset .
// @router /api/data/v2/datasets [POST]
func CreateDataset(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.CreateDataset)
}

// UpdateDataset .
// @router /api/data/v2/datasets/:dataset_id [PATCH]
func UpdateDataset(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.UpdateDataset)
}

// DeleteDataset .
// @router /api/data/v2/datasets/:dataset_id [DELETE]
func DeleteDataset(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.DeleteDataset)
}

// GetDataset .
// @router /api/data/v2/datasets/:dataset_id [GET]
func GetDataset(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.GetDataset)
}

// BatchGetDatasets .
// @router /api/data/v2/datasets/batch_get [POST]
func BatchGetDatasets(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.BatchGetDatasets)
}

// ImportDataset .
// @router /api/data/v2/datasets/:dataset_id/import [POST]
func ImportDataset(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.ImportDataset)
}

// GetDatasetIOJob .
// @router /api/data/v2/dataset_io_jobs/:job_id [GET]
func GetDatasetIOJob(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.GetDatasetIOJob)
}

// CreateDatasetVersion .
// @router /api/data/v2/datasets/:dataset_id/versions [POST]
func CreateDatasetVersion(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.CreateDatasetVersion)
}

// ListDatasetVersions .
// @router /api/data/v2/datasets/:dataset_id/versions/list [POST]
func ListDatasetVersions(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.ListDatasetVersions)
}

// GetDatasetVersion .
// @router /api/data/v2/dataset_versions/:version_id [GET]
func GetDatasetVersion(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.GetDatasetVersion)
}

// GetDatasetSchema .
// @router /api/data/v2/datasets/:dataset_id/schema [GET]
func GetDatasetSchema(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.GetDatasetSchema)
}

// UpdateDatasetSchema .
// @router /api/data/v2/datasets/:dataset_id/schema [PUT]
func UpdateDatasetSchema(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.UpdateDatasetSchema)
}

// BatchCreateDatasetItems .
// @router /api/data/v2/datasets/:dataset_id/items/batch_create [POST]
func BatchCreateDatasetItems(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.BatchCreateDatasetItems)
}

// UpdateDatasetItem .
// @router /api/data/v2/datasets/:dataset_id/items/:item_id [PUT]
func UpdateDatasetItem(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.UpdateDatasetItem)
}

// DeleteDatasetItem .
// @router /api/data/v2/datasets/:dataset_id/items/:item_id [DELETE]
func DeleteDatasetItem(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.DeleteDatasetItem)
}

// BatchDeleteDatasetItems .
// @router /api/data/v2/datasets/:dataset_id/items/batch_delete [POST]
func BatchDeleteDatasetItems(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.BatchDeleteDatasetItems)
}

// ListDatasetItems .
// @router /api/data/v2/datasets/:dataset_id/items/list [POST]
func ListDatasetItems(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.ListDatasetItems)
}

// ListDatasetItemsByVersion .
// @router /api/data/v2/datasets/:dataset_id/versions/:version_id/items/list [POST]
func ListDatasetItemsByVersion(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.ListDatasetItemsByVersion)
}

// GetDatasetItem .
// @router /api/data/v2/datasets/:dataset_id/items/:item_id [GET]
func GetDatasetItem(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.GetDatasetItem)
}

// BatchGetDatasetItems .
// @router /api/data/v2/datasets/:dataset_id/items/batch_get [POST]
func BatchGetDatasetItems(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.BatchGetDatasetItems)
}

// BatchGetDatasetItemsByVersion .
// @router /api/data/v2/datasets/:dataset_id/versions/:version_id/items/batch_get [POST]
func BatchGetDatasetItemsByVersion(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.BatchGetDatasetItemsByVersion)
}

// ListDatasets .
// @router /api/data/v2/datasets/list [POST]
func ListDatasets(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.ListDatasets)
}

// ListDatasetIOJobs .
// @router /api/data/v2/datasets/:dataset_id/io_jobs [POST]
func ListDatasetIOJobs(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.ListDatasetIOJobs)
}

// BatchGetDatasetVersions .
// @router /api/data/v2/dataset_versions/batch_get [POST]
func BatchGetDatasetVersions(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.BatchGetDatasetVersions)
}

// ClearDatasetItem .
// @router /api/evaluation/v3/datasets/:dataset_id/clear [POST]
func ClearDatasetItem(ctx context.Context, c *app.RequestContext) {
	invokeAndRender(ctx, c, localDataSvc.ClearDatasetItem)
}
