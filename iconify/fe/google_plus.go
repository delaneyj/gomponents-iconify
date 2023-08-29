package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feGooglePlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGooglePlus1" fill="currentColor"><path id="feGooglePlus2" d="M8.364 11.455h6.009c.054.318.1.636.1 1.054c0 3.636-2.437 6.218-6.11 6.218A6.359 6.359 0 0 1 2 12.364A6.359 6.359 0 0 1 8.364 6c1.718 0 3.154.627 4.263 1.664L10.9 9.327c-.473-.454-1.3-.982-2.536-.982c-2.173 0-3.946 1.8-3.946 4.019c0 2.218 1.773 4.018 3.946 4.018c2.518 0 3.463-1.81 3.609-2.746h-3.61v-2.181Zm13.636 0v1.818h-1.818v1.818h-1.818v-1.818h-1.819v-1.818h1.819V9.636h1.818v1.819H22Z"/></g></g>`),
		g.Group(children),
	)
}