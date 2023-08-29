package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudLinkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.03 5.507a4 4 0 0 1 7.94 0a3.252 3.252 0 0 1 2.932 2.447A3.488 3.488 0 0 0 12.5 7h-4a3.5 3.5 0 0 0-3.163 5H4.25a3.25 3.25 0 0 1-.22-6.493ZM8.5 8a2.5 2.5 0 0 0 0 5H9a.5.5 0 0 0 0-1h-.5a1.5 1.5 0 0 1 0-3H9a.5.5 0 0 0 0-1h-.5ZM12 8a.5.5 0 0 0 0 1h.5a1.5 1.5 0 0 1 0 3H12a.5.5 0 0 0 0 1h.5a2.5 2.5 0 0 0 0-5H12Zm-4 2.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}