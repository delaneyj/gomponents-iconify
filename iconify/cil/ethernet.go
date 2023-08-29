package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ethernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m385.128 118.178l-25.26 19.644l91.864 118.122l-91.919 118.236l25.263 19.64l107.192-137.88l-107.14-137.762zm-265.314-.001L12.621 255.993l107.138 137.826l25.263-19.638l-91.866-118.182l91.918-118.176l-25.26-19.646zM160 240h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32z"/>`),
		g.Group(children),
	)
}