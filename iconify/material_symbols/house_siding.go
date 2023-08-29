package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseSiding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-8H2l10-9l10 9h-3v8h-2v-2H7v2H5ZM9.45 8h5.1L12 5.7L9.45 8ZM7 12h10v-1.8l-.225-.2h-9.55L7 10.2V12Zm0 4h10v-2H7v2Z"/>`),
		g.Group(children),
	)
}