package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11h2V7H9v4Zm3 0h2V7h-2v4Zm3 0h2V7h-2v4ZM6 22q-.825 0-1.413-.588T4 20V8l6-6h8q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Zm0-2h12V4h-7.15L6 8.85V20Zm0 0h12H6Z"/>`),
		g.Group(children),
	)
}