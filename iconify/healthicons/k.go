package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func K(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 23.172V12a2 2 0 1 0-4 0v24a2 2 0 1 0 4 0v-7.172l2-2l10.586 10.586a2 2 0 1 0 2.828-2.828L22.828 24l10.586-10.586a2 2 0 1 0-2.828-2.828L18 23.172Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}