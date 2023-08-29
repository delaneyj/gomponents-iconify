package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Svg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.5 6.5c-.444 0-.843.193-1.118.5h-2.968l2.099-2.099a1.5 1.5 0 1 0-1.414-1.414L9 5.586V2.618a1.5 1.5 0 1 0-2 0v2.968L4.901 3.487a1.5 1.5 0 1 0-1.414 1.414L5.586 7H2.618a1.5 1.5 0 1 0 0 2h2.968l-2.099 2.099a1.5 1.5 0 1 0 1.414 1.414L7 10.414v2.968a1.5 1.5 0 1 0 2 0v-2.968l2.099 2.099a1.5 1.5 0 1 0 1.414-1.414L10.414 9h2.968A1.5 1.5 0 1 0 14.5 6.5z"/>`),
		g.Group(children),
	)
}