package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lemon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.536 3.393c3.905 3.906 3.905 10.237 0 14.143c-3.906 3.905-10.237 3.905-14.143 0L17.536 3.393"/><path d="M5.868 15.06a6.5 6.5 0 0 0 9.193-9.192m-4.597 4.596l4.597 4.597m-4.597-4.597v6.364m0-6.364h6.364"/></g>`),
		g.Group(children),
	)
}