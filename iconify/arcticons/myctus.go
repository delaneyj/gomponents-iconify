package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myctus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M24.45 17v2.7a2 2 0 0 1-3.414 1.414"/><path d="M24.45 13.7V17a2 2 0 1 1-4 0v-3.3m-9.9 5.291V11l4 8l4-7.988V19m-2.7 15.317v.033a2.65 2.65 0 1 1-5.3 0v-2.7A2.65 2.65 0 0 1 13.2 29h0a2.65 2.65 0 0 1 2.65 2.65v.033"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.75 29h5.3m-2.65 8v-8"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M24.95 29v5.35a2.65 2.65 0 1 0 5.3 0V29m4.55 0l2.65-3.65h-12.5M33.5 42.5a9 9 0 0 1 9-9"/><path d="M37.754 40.062c.245.32.553.438.98.438h.593a.998.998 0 0 0 .998-.998v-.004a.998.998 0 0 0-.998-.998h-.653a.999.999 0 0 1-1-.999h0c0-.552.45-1 1.002-1h.589c.428 0 .736.118.98.438"/></g>`),
		g.Group(children),
	)
}