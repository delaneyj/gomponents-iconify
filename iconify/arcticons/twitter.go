package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.74 16.55v1c0 10.07-7.64 21.61-21.62 21.61c-4.13 0-8.17-1.19-11.62-3.45c.6.08 1.2.12 1.81.11c3.42 0 6.75-1.13 9.44-3.24a7.563 7.563 0 0 1-7.1-5.29c.47.1.96.15 1.44.15c.68 0 1.35-.09 2-.27A7.574 7.574 0 0 1 7 19.72v-.1c1.05.59 2.23.91 3.44.94A7.541 7.541 0 0 1 8.05 10.5c3.86 4.75 9.56 7.64 15.68 7.94a6.38 6.38 0 0 1-.21-1.74a7.546 7.546 0 0 1 7.27-7.82c2.24-.08 4.4.84 5.9 2.51c1.7-.34 3.33-.97 4.83-1.85a7.654 7.654 0 0 1-3.39 4.27c1.51-.21 2.98-.63 4.37-1.26c-1.01 1.55-2.28 2.9-3.76 4Zm-8.39-.08l2.5-2.51m-2.5 0l2.5 2.51"/>`),
		g.Group(children),
	)
}