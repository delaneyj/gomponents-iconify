package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NSixtyFourSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.293 1h4.414l1 1h.793A3.5 3.5 0 0 1 15 5.5v4.838a1.662 1.662 0 0 1-3.24.525L11.14 9h-.735a.5.5 0 0 0-.498.45l-.26 2.607a2.157 2.157 0 0 1-4.294 0l-.26-2.607A.5.5 0 0 0 4.595 9H3.86l-.62 1.863A1.662 1.662 0 0 1 0 10.338V5.5A3.5 3.5 0 0 1 3.5 2h.793l1-1ZM4 7V6H3V5h1V4h1v1h1v1H5v1H4Zm5-1h1V5H9v1Zm3 0v1h-1V6h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}