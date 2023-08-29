package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checkbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<rect width="41" height="41" x="1.5" y="1.5" fill="none" stroke="currentColor" stroke-width="3" rx="2.5"/><path fill="currentColor" d="m17.758 26.254l13.198-13.211l2.36 2.358l-15.557 15.557l-7.072-7.071l2.359-2.358l4.712 4.725Z"/>`),
		g.Group(children),
	)
}