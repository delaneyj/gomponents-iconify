package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeReaderModeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h7V6H4v12Zm9 0h7V6h-7v12Zm1-8h5V8.5h-5V10Zm0 2.5h5V11h-5v1.5Zm0 2.5h5v-1.5h-5V15ZM4 6v12V6Z"/>`),
		g.Group(children),
	)
}