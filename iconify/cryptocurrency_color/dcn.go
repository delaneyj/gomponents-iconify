package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dcn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#136485"/><path fill="#FFF" d="M10.436 31.006a16.008 16.008 0 0 1-5.604-3.548l.147-.257c2.388-3.773 4.533-7.678 6.148-11.85c1.713-4.425 3.084-8.967 4.39-13.527c.117-.407.256-.807.384-1.21c.138.158.188.305.23.454c.82 2.926 1.613 5.86 2.464 8.776c1.55 5.313 3.73 10.353 6.617 15.077c.337.55.91 1.472 1.72 2.762a15.988 15.988 0 0 1-6.035 3.554a4320.193 4320.193 0 0 0-5.002-15.17l-.154-.002c-1.166 3.277-2.934 8.257-5.305 14.941zM16.075.049h-.124L16 0l.075.049z"/></g>`),
		g.Group(children),
	)
}