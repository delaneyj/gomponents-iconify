package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vtho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#2A5284"/><g fill="#FFF"><path fill-opacity=".704" d="M20.272 5.058h3.023l-5.682 7.584h5.099l-12.18 14.3l3.94-10.104h-4.45l4.487-11.78z"/><path d="M20.272 5.058h.182l-5.107 8.713h4.67l-9.485 13.17l3.94-10.103h-4.45l4.487-11.78z"/></g></g>`),
		g.Group(children),
	)
}