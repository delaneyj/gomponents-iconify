package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func B(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 12a2 2 0 0 1 2-2h10a8 8 0 0 1 5.292 14A8 8 0 0 1 26 38H16a2 2 0 0 1-2-2V12Zm12 10a4 4 0 0 0 0-8h-8v8h8Zm-8 4h8a4 4 0 0 1 0 8h-8v-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}