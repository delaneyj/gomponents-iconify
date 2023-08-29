package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 0C2.34 0 1 1.34 1 3c0 2 3 5 3 5s3-3 3-5c0-1.66-1.34-3-3-3zm0 1a2 2 0 0 1 2 2c0 1.11-.89 2-2 2a2 2 0 1 1 0-4z"/>`),
		g.Group(children),
	)
}