package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPharmacyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2l2-6l-2-6V5h12.7l1.45-4l2.35.85L18.35 5H21v2l-2 6l2 6v2H3Zm8-4h2v-3h3v-2h-3V9h-2v3H8v2h3v3Zm-5.9 2h13.8l-2-6l2-6H5.1l2 6l-2 6Zm6.9-6Z"/>`),
		g.Group(children),
	)
}