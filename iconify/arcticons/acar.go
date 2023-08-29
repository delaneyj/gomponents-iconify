package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 33.75c0-10.77 8.73-19.5 19.5-19.5s19.5 8.73 19.5 19.5h-39Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.935 23.32l-8.269 2.298a3.163 3.163 0 0 0-3.09-.147a3.174 3.174 0 0 0 2.848 5.674a3.163 3.163 0 0 0 1.73-2.567l6.78-5.258Z"/>`),
		g.Group(children),
	)
}