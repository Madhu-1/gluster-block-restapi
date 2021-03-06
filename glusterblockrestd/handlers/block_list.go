package handlers

import (
	"net/http"

	"github.com/gluster/gluster-block-restapi/pkg/utils"

	"github.com/gorilla/mux"
)

func blockVolumesListHandler(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	hostVolume := p["hostvolume"]
	cmdArgs := []string{"list", hostVolume, "--json"}
	out, err := utils.ExecuteCommandOutput(glusterBlockCLI, cmdArgs...)
	if err != nil {
		utils.SendHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendHTTPResponse(w, http.StatusOK, out)
}

func init() {
	registerRoute(Route{
		"BlockVolumesList",
		"GET",
		"/v1/blockvolumes/{hostvolume}",
		blockVolumesListHandler,
	})
}
