package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormCut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m18 7.524l-7.857 6.286L18 7.524Zm0 8.38L10.143 9.62L18 15.905Zm-9.429-5.761a1.571 1.571 0 1 0 0-3.143a1.571 1.571 0 0 0 0 3.143Zm0 6.286a1.571 1.571 0 1 0 0-3.143a1.571 1.571 0 0 0 0 3.143Z"/>`),
		g.Group(children),
	)
}