package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uploadcare/uploadcare-go/file"
	"github.com/uploadcare/uploadcare-go/test/testenv"
)

func listFiles(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	list, err := r.File.List(ctx, file.ListParams{})
	assert.Equal(t, nil, err)
	count := 0
	for list.Next() && count < 10 {
		res, err := list.ReadResult()
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEqual(t, "", res.ID)
		r.Artifacts.Files = append(r.Artifacts.Files, res)
		count++
	}
}

func fileInfo(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	info, err := r.File.Info(ctx, r.Artifacts.Files[0].ID)
	assert.Equal(t, nil, err)
	assert.Equal(
		t,
		r.Artifacts.Files[0].OriginalFileName,
		info.OriginalFileName,
	)
}

func fileStore(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	info, err := r.File.Store(ctx, r.Artifacts.Files[0].ID)
	assert.Equal(t, nil, err)
	assert.NotNil(t, info.StoredAt)
}

func fileDelete(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	info, err := r.File.Delete(ctx, r.Artifacts.Files[0].ID)
	assert.Equal(t, nil, err)
	assert.NotNil(t, info.RemovedAt)
}

func fileBatchStore(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	res, err := r.File.BatchStore(ctx, []string{r.Artifacts.Files[0].ID})
	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, len(res.Results))
}

func fileBatchDelete(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	ids := make([]string, 0, len(r.Artifacts.Files))
	for _, r := range r.Artifacts.Files {
		ids = append(ids, r.ID)
	}
	res, err := r.File.BatchDelete(ctx, ids)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, len(res.Results))
}

func fileCopy(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	_, err := r.File.Copy(
		ctx,
		file.CopyParams{
			LocalCopyParams: file.LocalCopyParams{
				Source: r.Artifacts.Files[0].ID,
			},
		},
	)
	assert.Equal(t, nil, err)
}

func fileLocalCopy(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	_, err := r.File.LocalCopy(
		ctx,
		file.LocalCopyParams{
			Source: r.Artifacts.Files[0].ID,
		},
	)
	assert.Equal(t, nil, err)
}

func fileRemoteCopy(t *testing.T, r *testenv.Runner) {
	ctx := context.Background()
	_, err := r.File.RemoteCopy(
		ctx,
		file.RemoteCopyParams{
			Source: r.Artifacts.Files[0].ID,
			Target: r.Artifacts.CustomStorage,
		},
	)
	assert.Equal(t, nil, err)
}
