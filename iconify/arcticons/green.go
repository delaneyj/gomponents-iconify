package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Green(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.586 9.88h26.897v18.827a12.132 12.132 0 0 1-5.38 9.414L28.034 43.5l-8.068-5.38a12.132 12.132 0 0 1-5.38-9.413Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.897 32.741a12.132 12.132 0 0 1-5.38-9.413V4.5h26.897"/>`),
		g.Group(children),
	)
}