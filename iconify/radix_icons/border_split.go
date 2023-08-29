package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor"><rect width="1" height="1" x="7" y="5.025" rx=".5"/><rect width="1" height="1" x="7" y="3.025" rx=".5"/><rect width="1" height="1" x="7" y="7.025" rx=".5"/><rect width="1" height="1" x="7" y="13.025" rx=".5"/><rect width="1" height="1" x="7" y="1.025" rx=".5"/><rect width="1" height="1" x="13" y="7.025" rx=".5"/><rect width="1" height="1" x="5" y="7.025" rx=".5"/><rect width="1" height="1" x="3" y="7.025" rx=".5"/><rect width="1" height="1" x="9" y="7.025" rx=".5"/><rect width="1" height="1" x="11" y="7.025" rx=".5"/><rect width="1" height="1" x="7" y="9.025" rx=".5"/><rect width="1" height="1" x="7" y="11.025" rx=".5"/><rect width="1" height="1" x="1" y="7.025" rx=".5"/><path fill-rule="evenodd" d="M1 1.5a.5.5 0 0 1 .5-.5H6v1H2v4H1V1.5ZM13 2H9V1h4.5a.5.5 0 0 1 .5.5V6h-1V2ZM1 13.5V9h1v4h4v1H1.5a.5.5 0 0 1-.5-.5Zm12-.5V9h1v4.5a.5.5 0 0 1-.5.5h-4v-1H13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}