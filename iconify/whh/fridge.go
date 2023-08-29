package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 960h-64v32q0 13-9.5 22.5T480 1024t-22.5-9.5T448 992v-32H192v32q0 13-9.5 22.5T160 1024t-22.5-9.5T128 992v-32H64q-26 0-45-19T0 896V448q0-27 18.5-45.5T64 384h512q27 0 45.5 18.5T640 448v448q0 26-18.5 45T576 960zM128 480q0-13-9.5-22.5T96 448t-22.5 9.5T64 480v64q0 13 9.5 22.5T96 576t22.5-9.5T128 544v-64zm448-160H64q-26 0-45-19T0 256V64q0-27 19-45.5T64 0h512q27 0 45.5 18.5T640 64v192q0 26-18.5 45T576 320zM128 160q0-13-9.5-22.5T96 128t-22.5 9.5T64 160v64q0 13 9.5 22.5T96 256t22.5-9.5T128 224v-64z"/>`),
		g.Group(children),
	)
}