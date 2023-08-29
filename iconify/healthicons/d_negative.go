package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsDNegative0)"><path d="M16 34V14h6c5.523 0 10 4.477 10 10s-4.477 10-10 10h-6Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM14 10a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h8c7.732 0 14-6.268 14-14s-6.268-14-14-14h-8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsDNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}