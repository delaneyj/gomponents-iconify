package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircledCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.68 5.5a1 1 0 0 1 1-1h6.643a1 1 0 0 1 1 1v3.875a1 1 0 0 1-1 1H9.68a1 1 0 0 1-1-1V5.5Zm2 1v1.875h4.643V6.5H10.68Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 19V8h1V6H7a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h6.337a5.526 5.526 0 0 1-1.737-2H8Zm10-7.793c.742.21 1.421.572 2 1.05V7a1 1 0 0 0-1-1h-2v2h1v3.207Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M16.5 20a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0 2a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M18.87 14.72a1 1 0 0 1 .156 1.405l-1.65 2.063a1.5 1.5 0 0 1-2.233.124l-1.104-1.105a1 1 0 0 1 1.414-1.414l.71.71l1.302-1.628a1 1 0 0 1 1.405-.156Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}