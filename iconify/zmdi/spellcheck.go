package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spellcheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m213 277l-24-64H68l-23 64H0L109 0h40l109 277h-45zM85 171h88L129 53zm323 12l30 30l-202 203l-109-109l30-30l79 79z"/>`),
		g.Group(children),
	)
}