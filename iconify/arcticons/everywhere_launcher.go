package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EverywhereLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="24" r="21.5"/><path d="M2.586 22.834H26.73v2.945H2.627m1.812-10.095h22.294v2.945H3.363M9.886 7.921h16.836v2.945H7.105M3.72 30.94h23.002v2.945H4.972M8.64 38.73h18.087v2.945H11.98m16.8-33.763v2.957m2.07-2.96v2.957m-2.066 4.817v2.958m2.069-2.961v2.958m-2.076 4.203v2.957m2.07-2.959v2.957m-2.059 5.138v2.958m2.07-2.96v2.957m-2.074 4.846v2.958m2.069-2.96v2.957"/></g>`),
		g.Group(children),
	)
}