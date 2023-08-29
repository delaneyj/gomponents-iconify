package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 35H4C4 27 9 17 14 17L16 26"/><path d="M28 35.0001C28 30 30 26.9893 37 27"/><path d="M44 28.5374C45 33.511 42 34.9999 40 33.9999C38 32.9999 38.5 28.0004 37 24C33.8599 15.6254 22 13.9997 15 21.0003"/></g>`),
		g.Group(children),
	)
}