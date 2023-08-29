package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor"><rect width="1" height="1" x="7" y="5.025" rx=".5"/><rect width="1" height="1" x="13" y="5.025" rx=".5"/><rect width="1" height="1" x="7" y="3.025" rx=".5"/><rect width="1" height="1" x="13" y="3.025" rx=".5"/><rect width="1" height="1" x="7" y="7.025" rx=".5"/><rect width="1" height="1" x="7" y="13.025" rx=".5"/><rect width="1" height="1" x="7" y="1.025" rx=".5"/><rect width="1" height="1" x="13" y="7.025" rx=".5"/><rect width="1" height="1" x="13" y="13.025" rx=".5"/><rect width="1" height="1" x="13" y="1.025" rx=".5"/><rect width="1" height="1" x="5" y="7.025" rx=".5"/><rect width="1" height="1" x="5" y="13.025" rx=".5"/><rect width="1" height="1" x="5" y="1.025" rx=".5"/><rect width="1" height="1" x="3" y="7.025" rx=".5"/><rect width="1" height="1" x="3" y="13.025" rx=".5"/><rect width="1" height="1" x="3" y="1.025" rx=".5"/><rect width="1" height="1" x="9" y="7.025" rx=".5"/><rect width="1" height="1" x="9" y="13.025" rx=".5"/><rect width="1" height="1" x="9" y="1.025" rx=".5"/><rect width="1" height="1" x="11" y="7.025" rx=".5"/><rect width="1" height="1" x="11" y="13.025" rx=".5"/><rect width="1" height="1" x="11" y="1.025" rx=".5"/><rect width="1" height="1" x="7" y="9.025" rx=".5"/><rect width="1" height="1" x="13" y="9.025" rx=".5"/><rect width="1" height="1" x="7" y="11.025" rx=".5"/><rect width="1" height="1" x="13" y="11.025" rx=".5"/><rect width="1" height="1" x="1" y="5.025" rx=".5"/><rect width="1" height="1" x="1" y="3.025" rx=".5"/><rect width="1" height="1" x="1" y="7.025" rx=".5"/><rect width="1" height="1" x="1" y="13.025" rx=".5"/><rect width="1" height="1" x="1" y="1.025" rx=".5"/><rect width="1" height="1" x="1" y="9.025" rx=".5"/><rect width="1" height="1" x="1" y="11.025" rx=".5"/></g>`),
		g.Group(children),
	)
}