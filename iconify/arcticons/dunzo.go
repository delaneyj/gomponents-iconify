package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dunzo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.703 19.327l5.809 4.209l6.436-16.233C29.744 6.16 27.18 5.5 24.339 5.5H15.46c-.673 0-1.249.483-1.366 1.146L8.163 40.284c-.032.18-.029.355-.01.526l11.55-21.483Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.944 29.523l-5.543-3.794L8.154 40.822a1.883 1.883 0 0 0 1.869 1.678h7.792c10.218 0 19.96-8.283 21.762-18.5h0c.533-3.022.295-5.866-.545-8.385L26.944 29.523Z"/>`),
		g.Group(children),
	)
}