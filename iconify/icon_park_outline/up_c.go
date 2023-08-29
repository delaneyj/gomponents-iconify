package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipOUpC0" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z" clip-rule="evenodd"/></defs><g fill="none" stroke="currentColor" stroke-width="4"><use href="#ipOUpC0" clip-rule="evenodd"/><use href="#ipOUpC0" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="m33 27l-9-9l-9 9"/></g>`),
		g.Group(children),
	)
}