package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomenCoat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 10c4-1 7-6 7-6h10s3 5.2 7 6l6 21h-8v13H14V31H6l6-21Z"/><path d="m19 4l5 14l5-14m-5 14v26"/></g>`),
		g.Group(children),
	)
}