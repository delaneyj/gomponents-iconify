package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func City(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 8v8h3V8H1zm2 7H1.979v-2.021H3V15zm.021-3H1.98V9.979h1.041V12zM10 5V3H9V0H8v3H7v2H5v11h7V5h-2zM7 15H6v-2h1v2zm0-2.958H6V10h1v2.042zm0-3H6V7h1v2.042zM9 15H8v-2h1v2zm0-2.958H8V10h1v2.042zm0-3H8V7h1v2.042zM11 15h-1v-2h1v2zm0-2.958h-1V10h1v2.042zm0-3h-1V7h1v2.042zM13 7v9h4v-5l-4-4zm2.031 8.062H14v-1.094h1.031v1.094zm0-2H14v-1.094h1.031v1.094zm0-2H14V9.968h1.031v1.094z"/>`),
		g.Group(children),
	)
}