package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ONegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsONegative0)"><path d="M14 24c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 10c-7.732 0-14 6.268-14 14s6.268 14 14 14s14-6.268 14-14s-6.268-14-14-14Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsONegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}