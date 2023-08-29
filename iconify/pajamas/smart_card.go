package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 3.5H2a.5.5 0 0 0-.5.5v8a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V4a.5.5 0 0 0-.5-.5ZM2 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H2Zm4 7h3v1a1 1 0 0 1-1 1H6V9Zm-2 2h1V9H3v1a1 1 0 0 0 1 1Zm3-6h1a1 1 0 0 1 1 1v2H7V5ZM4 5a1 1 0 0 0-1 1v2h3V5H4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}