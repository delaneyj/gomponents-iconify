package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m623.352 238l116 106c14 13 10 23-9 23h-87v277c0 19-16 36-35 36h-151V486c0-19-17-36-36-36h-95c-19 0-36 17-36 36v194h-151c-19 0-36-17-36-36V367h-86c-19 0-23-10-9-23l339-308c14-13 38-13 53 0l77 69V50c0-19 17-36 36-36h75c19 0 35 17 35 36v188z"/>`),
		g.Group(children),
	)
}