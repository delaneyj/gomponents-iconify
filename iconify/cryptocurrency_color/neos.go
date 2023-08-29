package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#e5f300"/><path fill="#fff" d="m10.5 9.358l8.143 4.926v3.097l-5.714-3.433V26H10.5zm11 13.284l-8.143-4.926V14.62l5.714 3.433V6H21.5z"/></g>`),
		g.Group(children),
	)
}