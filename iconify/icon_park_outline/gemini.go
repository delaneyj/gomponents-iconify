package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gemini(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 6s11.588 8 20 8s20-8 20-8M4 42s11.588-8 20-8s20 8 20 8M15 12v24m18-24v24"/>`),
		g.Group(children),
	)
}