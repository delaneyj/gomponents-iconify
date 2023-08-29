package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationDot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.617 8.677a4.5 4.5 0 1 0-7.235 0L8 13.5l3.617-4.823Zm1.203.897a6 6 0 1 0-9.64 0L6.875 14.5H4.75a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 0-1.5H9.125l3.695-4.926ZM8 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}