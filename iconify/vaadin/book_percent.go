package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookPercent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.6 2.5C11 1.3 11 0 11 0H2v12.5C2 14.4 3.6 16 5.5 16H14V3s-1-.2-1.4-.5zm-7.1.7c.8 0 1.5.7 1.5 1.6s-.7 1.4-1.5 1.4S4 5.6 4 4.8s.7-1.6 1.5-1.6zM9 3h1l-5 7H4l5-7zm1 5.5c0 .8-.7 1.5-1.5 1.5S7 9.3 7 8.5S7.7 7 8.5 7s1.5.7 1.5 1.5zm3 6.5H5.5c-1 0-1.8-.6-2-1.3c-.1-.4 0-.7.4-.7H11V2.7c0 .6 1 1.1 2 1.3v11z"/><path fill="currentColor" d="M9 8.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0zM6 4.8a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0z"/>`),
		g.Group(children),
	)
}