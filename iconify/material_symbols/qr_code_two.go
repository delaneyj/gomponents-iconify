package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21v-2h2v2h-2Zm-2-2v-5h2v5h-2Zm8-3v-4h2v4h-2Zm-2-4v-2h2v2h-2ZM5 14v-2h2v2H5Zm-2-2v-2h2v2H3Zm9-7V3h2v2h-2ZM4.5 7.5h3v-3h-3v3ZM3 9V3h6v6H3Zm1.5 10.5h3v-3h-3v3ZM3 21v-6h6v6H3ZM16.5 7.5h3v-3h-3v3ZM15 9V3h6v6h-6Zm2 12v-3h-2v-2h4v3h2v2h-4Zm-4-7v-2h4v2h-4Zm-4 0v-2H7v-2h6v2h-2v2H9Zm1-5V5h2v2h2v2h-4ZM5.25 6.75v-1.5h1.5v1.5h-1.5Zm0 12v-1.5h1.5v1.5h-1.5Zm12-12v-1.5h1.5v1.5h-1.5Z"/>`),
		g.Group(children),
	)
}