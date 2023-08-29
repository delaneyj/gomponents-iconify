package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M14 2a5 5 0 0 1 5 5v10a5 5 0 0 1-5 5h-4a5 5 0 0 1-5-5V7a5 5 0 0 1 5-5zm-2 4a1 1 0 0 0-1 1v4l.007.117A1 1 0 0 0 13 11V7l-.007-.117A1 1 0 0 0 12 6z"/></g>`),
		g.Group(children),
	)
}