package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 2.134a1 1 0 0 1 1 0l7.794 4.5a1 1 0 0 1 .5.866v7.845a3 3 0 0 1-1.5 2.598L13.5 21.29a3 3 0 0 1-3 0l-5.794-3.346a3 3 0 0 1-1.5-2.598V7.5a1 1 0 0 1 .5-.866l7.794-4.5ZM12 6a1 1 0 0 1 1 1v4.423l3.83 2.211a1 1 0 1 1-1 1.732L12 13.155l-3.83 2.211a1 1 0 1 1-1-1.732L11 11.423V7a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}