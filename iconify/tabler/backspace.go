package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H9l-5-5a1.5 1.5 0 0 1 0-2l5-5zm-8 4l4 4m0-4l-4 4"/>`),
		g.Group(children),
	)
}