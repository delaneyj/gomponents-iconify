package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bsd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#000"/><path fill="#FFF" d="M28 15.127H4l13.964-4.69L19.927 4L28 15.127zM4.11 16.655h23.78l-13.963 4.581l-1.963 6.655l-7.855-11.236z"/></g>`),
		g.Group(children),
	)
}