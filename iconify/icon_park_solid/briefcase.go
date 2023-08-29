package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBriefcase0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M32 16c0-6.075-3.582-12-8-12s-8 5.925-8 12"/><path fill="#fff" d="M9 16h30l1 12H27v-3h-6v3H8l1-12Z"/><path d="M8 28L6 42h36l-2-14"/><path d="M21 25h6v6h-6v-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBriefcase0)"/>`),
		g.Group(children),
	)
}