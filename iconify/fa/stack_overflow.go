package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1289 1632H171v-480H11v640h1438v-640h-160v480zm-942-524l33-157l783 165l-33 156zm103-374l67-146l725 339l-67 145zm201-356l102-123l614 513l-102 123zM1048 0l477 641l-128 96L920 96zM330 1471v-159h800v159H330z"/>`),
		g.Group(children),
	)
}