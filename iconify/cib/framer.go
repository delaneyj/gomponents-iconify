package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Framer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.333 0h21.333v10.667H15.999zm0 10.667H16l10.667 10.667H5.334zm0 10.666H16V32z"/>`),
		g.Group(children),
	)
}