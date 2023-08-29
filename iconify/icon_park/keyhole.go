package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyhole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z" clip-rule="evenodd"/><path fill="#43CCF8" stroke="#fff" d="M24 15C21.2386 15 19 17.2386 19 20C19 21.6358 19.7856 23.0882 21 24.0004L20 32H28L27.0005 24C28.2147 23.0878 29 21.6356 29 20C29 17.2386 26.7614 15 24 15Z"/></g>`),
		g.Group(children),
	)
}