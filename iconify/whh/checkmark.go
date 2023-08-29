package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1006 195L387 813q-18 18-53 17q-32-1-49-17L18 546Q0 528 0 502t18-44l89-89q18-18 44-18t44 18l141 141L829 18q18-18 44-18t44 18l89 88q18 19 18 44.5t-18 44.5z"/>`),
		g.Group(children),
	)
}