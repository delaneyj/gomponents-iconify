package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024H64q-27 0-45.5-18.5T0 960V64q0-26 18.5-45T64 0h448q106 0 181 75t75 181v512q0 106-75 181t-181 75zm128-768q0-53-37.5-90.5T512 128H160q-13 0-22.5 9.5T128 160v704q0 13 9.5 22.5T160 896h352q53 0 90.5-37.5T640 768V256z"/>`),
		g.Group(children),
	)
}