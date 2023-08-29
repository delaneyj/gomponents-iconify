package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardWireless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 2H8v2H6v2h2V4h8v2h2V4h-2V2zM8 8h2v2H8V8zm6 0V6h-4v2h4zm0 0h2v2h-2V8zM4 11h16v12H4V11zm14 10v-3H6v3h12zm0-6v-2H6v2h12z"/>`),
		g.Group(children),
	)
}