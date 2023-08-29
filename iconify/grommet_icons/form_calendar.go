package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M6 19h12V8H6v11Zm9-11V5v3ZM9 8V5v3Zm-3 3.5h12H6Z"/>`),
		g.Group(children),
	)
}