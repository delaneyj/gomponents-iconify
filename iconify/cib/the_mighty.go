package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheMighty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.573 0h-5.912l-3.593 18.651L12.412 0H6.473c.287.292.448.849.521 1.432v28.276c0 .828-.14 1.844-.568 2.292c1.355-.287 3.079-.552 4.552-.812V12.5L14.297 29c.093.615.181 1.448-.027 1.599c1.131-.172 2.26-.339 3.391-.489c-.14-.213-.083-.869-.005-1.416l3.271-16.496v17.464c1.407-.167 3.156-.323 4.625-.468l.021-.011c-.433-.437-.579-1.464-.579-2.297V2.303c0-.839.141-1.865.579-2.303z"/>`),
		g.Group(children),
	)
}