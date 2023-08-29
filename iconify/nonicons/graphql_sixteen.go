package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphqlSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M2.5 5.16a.5.5 0 0 1 .252-.433l5-2.858a.5.5 0 0 1 .496 0l5 2.858a.5.5 0 0 1 .252.434v5.678a.5.5 0 0 1-.252.434l-5 2.858a.5.5 0 0 1-.496 0l-5-2.857a.5.5 0 0 1-.252-.435V5.161ZM1 10.84a2 2 0 0 0 1.008 1.736l5 2.857a2 2 0 0 0 1.984 0l5-2.857A2 2 0 0 0 15 10.839V5.161a2 2 0 0 0-1.008-1.737l-5-2.857a2 2 0 0 0-1.984 0l-5 2.857A2 2 0 0 0 1 5.161v5.678Z"/><path d="M8 .27L1.182 11.75h13.636L8 .27Zm-4.182 9.98L8 3.208l4.182 7.042H3.818Z"/></g>`),
		g.Group(children),
	)
}