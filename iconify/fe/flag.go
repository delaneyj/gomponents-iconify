package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFlag0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFlag1" fill="currentColor"><path id="feFlag2" d="M6 12v9a1 1 0 0 1-2 0V4c0-1.1.9-2 2-2h6c1.1 0 1.999.9 1.999 2H18a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-5a2 2 0 0 1-1.997-2H6Zm0-8v6h7v2h5.002V6H12V4H6Z"/></g></g>`),
		g.Group(children),
	)
}