package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0H9.294c-.461 0-.692 0-.882-.082a1.002 1.002 0 0 1-.418-.337c-.12-.167-.167-.39-.261-.83L5.27 4.264c-.096-.451-.145-.677-.265-.845a1.003 1.003 0 0 0-.419-.338C4.397 3 4.167 3 3.707 3H3m3 3h12.873c.722 0 1.082 0 1.325.15a1 1 0 0 1 .435.579c.077.274-.022.621-.222 1.314l-1.385 4.8c-.12.415-.18.622-.3.776a1.004 1.004 0 0 1-.409.307c-.18.074-.396.074-.825.074H7.73M8 21a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}