package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nextdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31.99 13.089L16 3.297l-5.495 3.365V3.297H5.5v6.427L.005 13.089L2.62 17.36l2.88-1.755v13.099h21V15.605l2.875 1.755l2.615-4.271z"/>`),
		g.Group(children),
	)
}