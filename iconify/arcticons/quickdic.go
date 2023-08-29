package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quickdic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.4 6.45v35.1a2 2 0 0 0 1.95 2h2.38V4.5h-2.38A2 2 0 0 0 8.4 6.45Zm4.33-1.95v39h24.92a2 2 0 0 0 2-2V6.45a2 2 0 0 0-2-1.95Zm5.43 25.09l3.76-11.21m3.6 11.24l-3.6-11.24m2.39 7.48h-4.9"/><rect width="5.58" height="7.43" x="28.58" y="22.19" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.79"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.16 22.19v7.43"/>`),
		g.Group(children),
	)
}