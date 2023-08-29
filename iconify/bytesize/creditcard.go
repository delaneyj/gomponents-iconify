package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creditcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7v18h28V7Zm3 11h4m-4 3h6"/><path fill="currentColor" d="M2 11v2h28v-2Z"/></g>`),
		g.Group(children),
	)
}