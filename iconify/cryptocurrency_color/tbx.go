package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tbx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#5244D4"/><path fill="#FFF" d="M15.7 27.4C9.238 27.4 4 22.162 4 15.7C4 9.238 9.238 4 15.7 4c6.462 0 11.7 5.238 11.7 11.7c0 6.462-5.238 11.7-11.7 11.7zm2.89-7.7l-2.89-4l-2.89 4l-2.35-4l2.62-4.48h5.24l2.62 4.48l-2.35 4zm.86-10.4h-7.5l-3.74 6.4l3.75 6.4h7.49l3.74-6.4l-3.74-6.4z"/></g>`),
		g.Group(children),
	)
}