package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.338 15.934l-7.732-2.779l-3.232-4.058A2.99 2.99 0 0 0 16.054 8H8.058a2.998 2.998 0 0 0-2.48 1.312l-2.712 3.983A4.988 4.988 0 0 0 2 16.107V24a1 1 0 0 0 1 1h2.142a3.98 3.98 0 0 0 7.716 0h6.284a3.98 3.98 0 0 0 7.716 0H29a1 1 0 0 0 1-1v-7.125a1 1 0 0 0-.662-.941ZM9 26a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2Zm14 0a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm5-3h-1.142a3.98 3.98 0 0 0-7.716 0h-6.284a3.98 3.98 0 0 0-7.716 0H4v-6.893a2.998 2.998 0 0 1 .52-1.688l2.711-3.981A1 1 0 0 1 8.058 10h7.996a.993.993 0 0 1 .764.354l3.4 4.269a1 1 0 0 0 .444.318L28 17.578Z"/>`),
		g.Group(children),
	)
}