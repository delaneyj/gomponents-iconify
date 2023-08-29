package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ambulance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m2.264 18.689l5.557-7.145A4 4 0 0 1 10 10.121V10a1 1 0 0 1 1-1a1 1 0 1 1 2 0c.136 0 .265.027.383.076A5.001 5.001 0 0 1 18 6h9a4 4 0 0 1 4 4v13a4 4 0 0 1-4 4h-.035a3.501 3.501 0 0 1-6.93 0h-8.07a3.501 3.501 0 0 1-6.93 0H5a4 4 0 0 1-4-4v-.628a6 6 0 0 1 1.264-3.683ZM27 25a2 2 0 0 0 1.732-1H28a1 1 0 1 1 0-2h1v-1H15v4h5.337a3.5 3.5 0 0 1 6.326 0H27Zm-13 0V12h-3.022a2 2 0 0 0-1.579.772L9.222 13H12a1 1 0 0 1 1 1v2a2 2 0 0 1-2 2H5.333l-1.49 1.917a4 4 0 0 0-.6 1.083H4a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-.732A2 2 0 0 0 5 25h.337a3.5 3.5 0 0 1 6.326 0H14ZM28.732 9A2 2 0 0 0 27 8h-9a2.99 2.99 0 0 0-2.236 1h12.968ZM22 12a1 1 0 0 0-1 1v1h-1a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0v-1h1a1 1 0 1 0 0-2h-1v-1a1 1 0 0 0-1-1ZM10 26.5a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm15 0a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Z"/>`),
		g.Group(children),
	)
}