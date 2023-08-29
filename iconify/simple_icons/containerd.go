package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Containerd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.629 0v24H20.37V0zM17.59 21.208H6.421V10.604h7.812V6.692h3.346v14.516zm-7.823-7.812h4.466v5.02H9.767z"/>`),
		g.Group(children),
	)
}