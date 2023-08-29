package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircledOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.68 2.5a1 1 0 0 1 1-1h6.643a1 1 0 0 1 1 1v3.875a1 1 0 0 1-1 1H6.68a1 1 0 0 1-1-1V2.5Zm2 1v1.875h4.643V3.5H7.68Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5 16V5h1V3H4a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h6.337A5.526 5.526 0 0 1 8.6 16H5Zm10-7.793c.742.21 1.421.572 2 1.05V4a1 1 0 0 0-1-1h-2v2h1v3.207Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.5 17a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0 2a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.87 11.72a1 1 0 0 1 .156 1.405l-1.65 2.063a1.5 1.5 0 0 1-2.232.124l-1.105-1.105a1 1 0 0 1 1.414-1.414l.71.71l1.302-1.628a1 1 0 0 1 1.405-.156Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}