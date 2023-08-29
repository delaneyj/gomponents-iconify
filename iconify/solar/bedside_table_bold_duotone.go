package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedsideTableBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 10c0-.442.002-1.608.004-2H22c.002.392 0 1.558 0 2v2c0 .442 0 1.608-.002 2H2.002C2 13.608 2 12.442 2 12v-2Z" clip-rule="evenodd" opacity=".5"/><path d="M13 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M2.026 7.25c.066-2.021.302-3.235 1.146-4.078C4.343 2 6.229 2 10 2h4c3.771 0 5.657 0 6.829 1.172c.843.843 1.08 2.057 1.145 4.078L22 8H2l.026-.75ZM12 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-9.974 8.75c.066 2.021.302 3.235 1.146 4.078a3.1 3.1 0 0 0 1.078.697V22a.75.75 0 0 0 1.5 0v-2.129C6.82 20 8.194 20 10 20h4c1.806 0 3.18 0 4.25-.129V22a.75.75 0 1 0 1.5 0v-2.475a3.1 3.1 0 0 0 1.079-.697c.843-.843 1.08-2.057 1.145-4.078V14H2l.026.75ZM12 18a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}