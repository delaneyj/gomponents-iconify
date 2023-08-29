package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Target(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M21.417 18.583a6 6 0 1 0 8 8"/><path d="M28.28 11.72C26.94 11.255 25.5 11 24 11c-7.18 0-13 5.82-13 13s5.82 13 13 13s13-5.82 13-13c0-1.5-.254-2.94-.72-4.28"/><path d="M33.568 6.433A19.911 19.911 0 0 0 24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20c0-3.466-.882-6.726-2.432-9.568M44 4L24 24"/></g>`),
		g.Group(children),
	)
}