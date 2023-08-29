package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arrowleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M689.13 466V251c0-26-19-45-45-45h-250V36c0-16-8-27-22-34c-5-1-10-2-14-2c-10 0-18 3-24 10l-324 324c-14 12-13 36 0 50l324 324c20 22 60 7 60-26V510h250c26 0 45-18 45-44z"/>`),
		g.Group(children),
	)
}