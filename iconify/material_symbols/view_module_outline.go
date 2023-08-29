package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewModuleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.025 19V5H21v14H3.025Zm12.65-8H19V7h-3.325v4Zm-5.325 0h3.325V7H10.35v4Zm-5.325 0H8.35V7H5.025v4Zm0 6H8.35v-4H5.025v4Zm5.325 0h3.325v-4H10.35v4Zm5.325 0H19v-4h-3.325v4Z"/>`),
		g.Group(children),
	)
}