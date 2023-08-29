package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escapepod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.68 42.89L5.5 31.7l9.7-4.52h0a10.15 10.15 0 0 0 6.14 5.39Zm6.16-14.67a3.4 3.4 0 1 1 3.39-3.41h0a3.4 3.4 0 0 1-3.39 3.41Zm-1.49-15a13 13 0 0 1 13.06 13v.1M21.35 5.11A21.15 21.15 0 0 1 42.5 26.25v.06"/>`),
		g.Group(children),
	)
}