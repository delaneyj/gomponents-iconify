package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorshipJewish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24.291 16l3.585-6.518A1 1 0 0 0 27 8h-7.109l-3.015-5.482a1 1 0 0 0-1.752 0L12.109 8H5a1 1 0 0 0-.876 1.482L7.709 16l-3.585 6.518A1 1 0 0 0 5 24h7.109l3.015 5.482a1 1 0 0 0 1.752 0L19.891 24H27a1 1 0 0 0 .876-1.482Zm-5.582 6L16 26.925L13.291 22h-6.6l3.3-6l-3.3-6h6.6L16 5.075L18.709 10h6.6l-3.3 6l3.3 6Z"/>`),
		g.Group(children),
	)
}