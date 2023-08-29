package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourPointConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 28v12h12m8 0h12V28m0-8V8H28m-8 0H8v12m36 0h-8v8h8v-8Zm-32 0H4v8h8v-8Zm16 16h-8v8h8v-8Zm0-32h-8v8h8V4Zm-4 14v12m-6-6h12M28 8l12 12M20 8L8 20m12 20L8 28m32 0L29 40"/>`),
		g.Group(children),
	)
}