package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eboks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.928 30.464a7.924 7.924 0 0 1-14.84-3.812l-.037-5.15a7.928 7.928 0 0 1 15.855-.115l.018 2.575l-15.855.115M43.027 34h-8"/>`),
		g.Group(children),
	)
}