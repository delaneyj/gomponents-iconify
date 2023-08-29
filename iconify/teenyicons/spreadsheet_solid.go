package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpreadsheetSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10 7.995V10H8V7.995h2Zm0-2.998v1.998H8V4.997h2Zm-3 0H5v1.998h2V4.997Zm0 2.998H5V10h2V7.995Z"/><path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h8.207L14 3.293V13.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12Zm10 2.497H4V11h7V3.997Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}