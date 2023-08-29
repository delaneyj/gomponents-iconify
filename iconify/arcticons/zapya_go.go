package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZapyaGo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.66 16.285a54.148 54.148 0 0 1 23.777-4.784s-6.814 9.762-7.54 15.32c0 0 13.822-2.368 21.603 0l-.677 2.561s-17.446-.628-29.817 7.153c-.725-6.476 3.044-17.35 3.044-17.35a49.369 49.369 0 0 0-11.55 1.45Z"/>`),
		g.Group(children),
	)
}