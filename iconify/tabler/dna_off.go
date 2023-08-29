package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnaOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 12a3.898 3.898 0 0 0-1.172-2.828A4.027 4.027 0 0 0 12 8M9.172 9.172a4 4 0 1 0 5.656 5.656"/><path d="M9.172 20.485a4 4 0 1 0-5.657-5.657M14.828 3.515a4 4 0 1 0 5.657 5.657M3 3l18 18"/></g>`),
		g.Group(children),
	)
}