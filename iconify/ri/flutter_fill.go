package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlutterFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.504 2.001l-10 10l3.083 3.083l13.08-13.083h-6.163Zm-.005 9.198l-5.376 5.42L13.496 22h6.188l-5.387-5.4l5.389-5.4h-6.188Z"/>`),
		g.Group(children),
	)
}