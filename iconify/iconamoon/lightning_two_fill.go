package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 1a1 1 0 0 0-.96.72l-3.5 12A1 1 0 0 0 6 15h3.867l-.86 6.876a1 1 0 0 0 1.825.679l8-12A1 1 0 0 0 18 9h-3.557l2.493-6.649A1 1 0 0 0 16 1H9.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}