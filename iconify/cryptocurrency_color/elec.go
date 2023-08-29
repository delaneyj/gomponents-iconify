package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elec(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#F90"/><g fill="#FFF"><path d="m10.76 27.587l12.666-13.303H15.76z"/><path d="M8 18.27h7.666l7.76-3.986H15.76z"/><path d="M19.51 4L8 18.27h7.666z"/></g></g>`),
		g.Group(children),
	)
}