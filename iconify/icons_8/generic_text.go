package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenericText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 6L8 26h2l2.094-6h7.812L22 26h2L17 6h-2zm1 2.844L19.188 18h-6.375L16 8.844z"/>`),
		g.Group(children),
	)
}