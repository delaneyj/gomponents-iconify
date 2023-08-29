package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeTwoAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 14v-2h2v2H5Zm-2-2v-2h2v2H3Zm9-7V3h2v2h-2ZM4.5 7.5h3v-3h-3v3ZM3 9V3h6v6H3Zm1.5 10.5h3v-3h-3v3ZM3 21v-6h6v6H3ZM16.5 7.5h3v-3h-3v3ZM15 9V3h6v6h-6Zm-6 5v-2H7v-2h4v4H9Zm1-5V5h2v2h2v2h-4ZM5.25 6.75v-1.5h1.5v1.5h-1.5Zm0 12v-1.5h1.5v1.5h-1.5Zm12-12v-1.5h1.5v1.5h-1.5ZM16 21v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}