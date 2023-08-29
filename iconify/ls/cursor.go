package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M298 312v42h42v42H213v85h42v85h43v85h-85v-85h-43v-85h-42v-42H85v42H43v43H0V57h43v42h42v43h43v42h42v43h43v42h42v43h43z"/>`),
		g.Group(children),
	)
}