package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyTwoHundredThreeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M24 18a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path fill-rule="evenodd" d="M15 16.447V34.04c0 3.412 2.813 7.96 9 7.96c6.188 0 9-3.98 9-7.959v-.855c-1.444-3.67-4.435-7.343-8.014-10.434a40.935 40.935 0 0 0-6.371-4.522a30.812 30.812 0 0 1-.568-.317A30.837 30.837 0 0 0 15 16.447Zm10.032 3.744c.427.34.847.69 1.261 1.048c2.6 2.245 4.969 4.839 6.707 7.592V17c-3.042 2.178-5.51 3.186-7.968 3.191Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}