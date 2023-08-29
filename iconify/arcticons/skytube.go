package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skytube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.25 38.49H11.07a6.56 6.56 0 0 1 0-13.11a6.46 6.46 0 0 1 3.13.78a7 7 0 0 1-.1-1a6.56 6.56 0 0 1 13.09-.76a5.08 5.08 0 0 1 0 .78a6.38 6.38 0 0 1-.09 1.08a6.56 6.56 0 0 1 6.4 11.45a6.71 6.71 0 0 1-3.22.83Zm-4.89-18.26v-10a.81.81 0 0 1 .84-.78a1 1 0 0 1 .34.09l16.56 9.6a.81.81 0 0 1 0 1.4l-9.77 5.62"/>`),
		g.Group(children),
	)
}