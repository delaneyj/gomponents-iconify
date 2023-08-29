package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1395 709L781 323l92-151h855zM373 974L189 858V0l1183 743zm1019-135l147 95v858l-532-335zm-37-21l-500 802H0l356-571z"/>`),
		g.Group(children),
	)
}