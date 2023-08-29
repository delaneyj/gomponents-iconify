package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.15l-2-2V16h-1.15l-2-2H20v-4h-4v3.15l-2-2V10h-1.15l-2-2H14V4h-4v3.15l-2-2V4H6.85l-2-2H20q.825 0 1.413.588T22 4v15.15ZM16 8h4V4h-4v4ZM4 22q-.825 0-1.413-.588T2 20V4.8l-.6-.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.3.725q-.275.275-.675.275t-.7-.275L19.15 22H4Zm12-2h1.15L16 18.8V20Zm-6-6h1.15L10 12.8V14Zm0 6h4v-3.2l-.85-.8H10v4ZM4 8h1.15L4 6.8V8Zm0 6h4v-3.2l-.85-.8H4v4Zm4 6v-4H4v4h4Z"/>`),
		g.Group(children),
	)
}