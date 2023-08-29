package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallBellBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M24 180h208a12 12 0 0 0 0-24h-4.09A100.16 100.16 0 0 0 140 60.72V44h12a12 12 0 0 0 0-24h-48a12 12 0 0 0 0 24h12v16.72A100.16 100.16 0 0 0 28.09 156H24a12 12 0 0 0 0 24Zm104-96a76.09 76.09 0 0 1 75.89 72H52.11A76.09 76.09 0 0 1 128 84Zm116 124a12 12 0 0 1-12 12H24a12 12 0 0 1 0-24h208a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}