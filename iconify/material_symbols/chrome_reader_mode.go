package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChromeReaderMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm9-2h7V6h-7v12Zm1-3h5v-1.5h-5V15Zm0-2.5h5V11h-5v1.5Zm0-2.5h5V8.5h-5V10Z"/>`),
		g.Group(children),
	)
}