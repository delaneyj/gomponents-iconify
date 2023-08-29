package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.113 15.196V8.804a1 1 0 0 0-1.53-.848l-5.113 3.196a1 1 0 0 0 0 1.696l5.113 3.196a1 1 0 0 0 1.53-.848Zm-8 0V8.804a1 1 0 0 0-1.53-.848L4.47 11.152a1 1 0 0 0 0 1.696l5.113 3.196a1 1 0 0 0 1.53-.848Z"/>`),
		g.Group(children),
	)
}