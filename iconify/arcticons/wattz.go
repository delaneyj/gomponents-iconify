package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wattz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5h-7.75a2 2 0 0 1-2-2V10.4a2 2 0 0 1 2-2h3.85V4.5h7.8v3.9h3.85a2 2 0 0 1 2 2v31.1a2 2 0 0 1-2 2H24Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M26 14.617v9.73h4l-8 13.27v-9.73h-4l8-13.27l-8 13.27h4v9.73l8-13.27h-4v-9.73"/>`),
		g.Group(children),
	)
}