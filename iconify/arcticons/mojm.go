package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mojm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 33.598v-8.586l6.04-6.04h7.432l9.74 9.74l-.388.388h-9.14l-.644.643l4.38 4.38l-8.376 8.377L5.5 33.598Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.598 5.5h-8.586l-6.04 6.04v7.432l9.74 9.74l.388-.388v-9.14l.643-.644l4.38 4.38l8.377-8.376L33.598 5.5Z"/>`),
		g.Group(children),
	)
}