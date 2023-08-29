package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Android(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M43.9011 36H4.09863C5.10208 25.8934 13.6292 18 23.9999 18C34.3706 18 42.8977 25.8934 43.9011 36Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 20L10 13"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M33 20L37 13"/><circle cx="15" cy="29" r="2" fill="#fff"/><circle cx="33" cy="29" r="2" fill="#fff"/></g>`),
		g.Group(children),
	)
}