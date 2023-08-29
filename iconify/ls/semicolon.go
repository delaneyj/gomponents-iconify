package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Semicolon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M67 145h77v77H67v-77zm17 393l70 33l-93 196l-61-29z"/>`),
		g.Group(children),
	)
}