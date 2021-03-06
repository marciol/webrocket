package main

func deleteVhost(params []string) (err error, ok bool) {
	var path string
	if path, ok = vhostParams(params); !ok {
		return
	}
	_, err = performRequest("DELETE", "/vhost", "", map[string]string{
		"path": path,
	})
	return
}
