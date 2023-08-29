package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M278 352H174l297 383h-91L83 352H71v383H0V0h278c58 0 110 27 142 71c22 30 34 66 34 105s-12 75-34 104c-32 44-84 72-142 72zm3-281H71v209h210c56-1 102-47 102-104c0-56-46-103-102-105z"/>`),
		g.Group(children),
	)
}