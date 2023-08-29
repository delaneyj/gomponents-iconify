package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveStarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 13.09V10h2v3.09c-.33-.05-.66-.09-1-.09c-.34 0-.67.04-1 .09M9.5 11c-.28 0-.5.22-.5.5V13h6v-1.5c0-.28-.22-.5-.5-.5h-5M6 19v-9H4v11h9.35c-.22-.63-.35-1.3-.35-2H6M21 9H3V3h18v6m-2-4H5v2h14V5m4 12.89l-2.89-.25L19 15l-1.13 2.64l-2.87.25l2.18 1.88l-.68 2.81l2.5-1.49l2.45 1.49l-.65-2.81l2.2-1.88Z"/>`),
		g.Group(children),
	)
}