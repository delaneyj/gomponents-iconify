package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LearnJavascript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.599 6.421V41.58L5.5 37.885m37-31.075H27.368v18.915l12.553-1.633l-1.032 14.014l-12.99 3.473"/>`),
		g.Group(children),
	)
}