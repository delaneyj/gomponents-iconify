package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wildberries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M34.06 24a3.713 3.713 0 1 1 0 7.426h-6.126V16.573h6.126a3.713 3.713 0 1 1 0 7.427Zm0-.001h-6.126m-2.855-7.424l-3.713 14.852l-3.713-14.852l-3.713 14.852l-3.713-14.852"/><path d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/></g>`),
		g.Group(children),
	)
}