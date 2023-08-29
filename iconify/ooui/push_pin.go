package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 8V2a2 2 0 0 0 2-2H5a2 2 0 0 0 2 2v6H6a2 2 0 0 0-2 2v1h5v5l1 4l1-4v-5h5v-1a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}