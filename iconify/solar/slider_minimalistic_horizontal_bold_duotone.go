package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SliderMinimalisticHorizontalBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M20.25 2.77a.76.76 0 0 1-.75.768h-15a.76.76 0 0 1-.75-.769A.76.76 0 0 1 4.5 2h15a.76.76 0 0 1 .75.77Zm0 18.46a.76.76 0 0 1-.75.77h-15a.76.76 0 0 1-.75-.77a.76.76 0 0 1 .75-.768h15a.76.76 0 0 1 .75.769Z" clip-rule="evenodd" opacity=".5"/><path d="M16 5.846c2.828 0 4.243 0 5.121.901C22 7.65 22 9.1 22 12c0 2.901 0 4.352-.879 5.253c-.878.9-2.293.9-5.121.9H8c-2.828 0-4.243 0-5.121-.9C2 16.352 2 14.9 2 12s0-4.351.879-5.253c.878-.9 2.293-.9 5.121-.9h8Z"/></g>`),
		g.Group(children),
	)
}