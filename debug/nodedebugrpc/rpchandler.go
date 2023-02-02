package nodedebugrpc

import (
	"context"
	"github.com/anytypeio/any-sync-node/debug/nodedebugrpc/nodedebugrpcproto"
	"github.com/anytypeio/any-sync-node/nodespace"
	"github.com/anytypeio/any-sync-node/nodestorage"
	"github.com/anytypeio/any-sync/commonspace/object/treegetter"
)

type rpcHandler struct {
	treeCache      treegetter.TreeGetter
	spaceService   nodespace.Service
	storageService nodestorage.NodeStorage
}

func (r *rpcHandler) DumpTree(ctx context.Context, request *nodedebugrpcproto.DumpTreeRequest) (resp *nodedebugrpcproto.DumpTreeResponse, err error) {
	tree, err := r.treeCache.GetTree(context.Background(), request.SpaceId, request.DocumentId)
	if err != nil {
		return
	}
	// TODO: commented
	_ = tree
	/*
		dump, err := tree.DebugDump(nil)
		if err != nil {
			return
		}*/
	resp = &nodedebugrpcproto.DumpTreeResponse{
		//Dump: dump,
	}
	return
}

func (r *rpcHandler) AllTrees(ctx context.Context, request *nodedebugrpcproto.AllTreesRequest) (resp *nodedebugrpcproto.AllTreesResponse, err error) {
	space, err := r.spaceService.GetSpace(ctx, request.SpaceId)
	if err != nil {
		return
	}
	heads := space.DebugAllHeads()
	var trees []*nodedebugrpcproto.Tree
	for _, head := range heads {
		trees = append(trees, &nodedebugrpcproto.Tree{
			Id:    head.Id,
			Heads: head.Heads,
		})
	}
	resp = &nodedebugrpcproto.AllTreesResponse{Trees: trees}
	return
}

func (r *rpcHandler) AllSpaces(ctx context.Context, request *nodedebugrpcproto.AllSpacesRequest) (resp *nodedebugrpcproto.AllSpacesResponse, err error) {
	ids, err := r.storageService.AllSpaceIds()
	if err != nil {
		return
	}
	resp = &nodedebugrpcproto.AllSpacesResponse{SpaceIds: ids}
	return
}

func (r *rpcHandler) TreeParams(ctx context.Context, request *nodedebugrpcproto.TreeParamsRequest) (resp *nodedebugrpcproto.TreeParamsResponse, err error) {
	tree, err := r.treeCache.GetTree(context.Background(), request.SpaceId, request.DocumentId)
	if err != nil {
		return
	}
	resp = &nodedebugrpcproto.TreeParamsResponse{
		RootId:  tree.Root().Id,
		HeadIds: tree.Heads(),
	}
	return
}
