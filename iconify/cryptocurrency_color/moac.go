package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#000"/><path fill="#FFF" d="M15.792 15.792L9.423 9.423l-.138-.208L7.02 7v17.515h2.284V12.4l4.916 4.985l1.592 1.592l1.592-1.592L22.32 12.4v12.115h2.284V7L22.32 9.215z"/></g>`),
		g.Group(children),
	)
}