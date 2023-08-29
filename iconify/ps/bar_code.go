package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0h43v427H0V0zm171 427V0H85v427h86zM213 0h43v427h-43V0zm86 0h42v427h-42V0zM0 469h512v43H0v-43zM469 0h-42v427h85V0h-43z"/>`),
		g.Group(children),
	)
}