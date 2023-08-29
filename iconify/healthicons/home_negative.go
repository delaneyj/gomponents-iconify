package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsHomeNegative0)"><path d="M34 25v11h-6V25h6Zm-9 0H14v6h11v-6Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM5 36a1 1 0 1 0 0 2h36a1 1 0 1 0 0-2h-3V22l-14-8l-14 8v14H5Zm31-19.101V11h-5v2.957l-6.493-3.819a1 1 0 0 0-1.014 0L5.631 20.645l1.014 1.724L24 12.16l17.355 10.21l1.014-1.724L36 16.9Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsHomeNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}