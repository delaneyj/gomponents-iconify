package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SofaThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4.75 17.75v2a.75.75 0 0 1-1.5 0v-2.145A2 2 0 0 1 4 13.75h16a2 2 0 0 1 .75 3.855v2.145a.75.75 0 0 1-1.5 0v-2H4.75Z"/><path d="M19.75 12.75c.966-3.884 1.395-5.903.384-7.322a4.002 4.002 0 0 0-.215-.275C18.72 3.75 16.48 3.75 11.999 3.75c-4.48 0-6.721 0-7.919 1.403c-.075.089-.147.18-.215.275c-1.011 1.42-.582 3.438.385 7.322l.2 1h15.1l.2-1Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}