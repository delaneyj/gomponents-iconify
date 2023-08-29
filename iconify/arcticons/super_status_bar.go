package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperStatusBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m9.365 24.028l.002-.016l-.002-11.039L24.042 4.5l14.593 8.425v10.958m0 5.952V35l-14.723 8.5l-14.547-8.241h0V30.17"/><path d="m14.107 15.783l9.74-5.624l9.953 5.746v5.197l-9.682-5.59l-4.793 2.768l14.52 8.383l4.746-2.773m.041 5.683L23.85 38.127l-14.462-8.56"/><path d="m9.435 24.12l14.672 8.47l4.767-2.752l-14.87-8.586v-5.271"/></g>`),
		g.Group(children),
	)
}