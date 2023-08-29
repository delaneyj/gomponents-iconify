package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crayon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.24 44.338V14.902A11.273 11.273 0 0 0 24 3.662h0a11.273 11.273 0 0 0-11.24 11.24v29.374Z"/><circle cx="19.611" cy="15.13" r="1.967" fill="none" stroke="currentColor" stroke-linejoin="round"/><circle cx="28.389" cy="15.13" r="1.967" fill="none" stroke="currentColor" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M26.49 20.166a.779.779 0 0 1 .708 1.113a3.581 3.581 0 0 1-6.398-.001a.779.779 0 0 1 .709-1.112Z"/>`),
		g.Group(children),
	)
}