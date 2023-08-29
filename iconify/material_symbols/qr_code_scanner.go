package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7V2h5v2H4v3H2Zm0 15v-5h2v3h3v2H2Zm15 0v-2h3v-3h2v5h-5Zm3-15V4h-3V2h5v5h-2Zm-2.5 10.5H19V19h-1.5v-1.5Zm0-3H19V16h-1.5v-1.5ZM16 16h1.5v1.5H16V16Zm-1.5 1.5H16V19h-1.5v-1.5ZM13 16h1.5v1.5H13V16Zm3-3h1.5v1.5H16V13Zm-1.5 1.5H16V16h-1.5v-1.5ZM13 13h1.5v1.5H13V13Zm6-8v6h-6V5h6Zm-8 8v6H5v-6h6Zm0-8v6H5V5h6ZM9.5 17.5v-3h-3v3h3Zm0-8v-3h-3v3h3Zm8 0v-3h-3v3h3Z"/>`),
		g.Group(children),
	)
}