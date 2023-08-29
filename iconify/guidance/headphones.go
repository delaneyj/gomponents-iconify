package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M22.5 13.5a6.575 6.575 0 0 1-4.369-1.66l-.381-.34h-.25v11h.25l.381-.34A6.575 6.575 0 0 1 22.5 20.5V12c0-5.799-4.701-10.5-10.5-10.5S1.5 6.201 1.5 12v8.5c1.61 0 3.165.591 4.369 1.66l.381.34h.25v-11h-.25l-.381.34A6.576 6.576 0 0 1 1.5 13.5"/>`),
		g.Group(children),
	)
}