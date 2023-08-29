package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.061 7.573A3.982 3.982 0 0 0 12 5c0-2.206-1.794-4-4-4H3v14h6c2.206 0 4-1.794 4-4a4.002 4.002 0 0 0-1.939-3.427zM6 3h1.586c.874 0 1.586.897 1.586 2s-.711 2-1.586 2H6V3zm2.484 10H6V9h2.484c.913 0 1.656.897 1.656 2s-.743 2-1.656 2z"/>`),
		g.Group(children),
	)
}