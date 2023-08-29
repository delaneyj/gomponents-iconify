package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobilityServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.338 19.934l-7.732-2.779l-3.232-4.058A2.99 2.99 0 0 0 16.054 12H8.058a2.998 2.998 0 0 0-2.48 1.312l-2.712 3.983A4.988 4.988 0 0 0 2 20.107V26a1 1 0 0 0 1 1h2.142a3.98 3.98 0 0 0 7.716 0h6.284a3.98 3.98 0 0 0 7.716 0H29a1 1 0 0 0 1-1v-5.125a1 1 0 0 0-.662-.941zM9 28a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zm14 0a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zm5-3h-1.142a3.98 3.98 0 0 0-7.716 0h-6.284a3.98 3.98 0 0 0-7.716 0H4v-4.893a2.998 2.998 0 0 1 .52-1.688l2.711-3.981A1 1 0 0 1 8.058 14h7.996a.993.993 0 0 1 .764.354l3.4 4.269a1 1 0 0 0 .444.318L28 21.578zm0-23h-4a2.002 2.002 0 0 0-2 2v8a2.002 2.002 0 0 0 2 2h4a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2zm0 2v2h-4V4zm-4 8V8h4v4z"/>`),
		g.Group(children),
	)
}