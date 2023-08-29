package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Historic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 1c-1 0-1 1-2 1H2c-.545 0-1 .455-1 1v7c0 .545.455 1 1 1h5v2.5s0 .5.5.5s.5-.5.5-.5V11h5c.545 0 1-.455 1-1V3c0-.545-.455-1-1-1H9.5c-1 0-1-1-2-1zM3 5V4h9v1zm0 1h4v1H3zm0 2h7v1H3z"/>`),
		g.Group(children),
	)
}