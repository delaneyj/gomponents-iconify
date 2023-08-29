package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShaverOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M36 14H12V6.02L16.474 4l4.245 2.02L24.193 4l3.474 2.02L31.912 4L36 6.02V14Zm-24 0v19c0 15 24 15 24 0V14M20 35h8"/><circle cx="24" cy="25" r="4"/></g>`),
		g.Group(children),
	)
}