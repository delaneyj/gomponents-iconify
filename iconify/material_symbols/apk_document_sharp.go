package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApkDocumentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19h10q-.1-1.225-.75-2.25t-1.7-1.625l.95-1.7q.05-.1.025-.225t-.15-.175q-.1-.05-.213-.025t-.162.125l-.975 1.75q-.5-.2-1-.313T12 14.45q-.525 0-1.025.113t-1 .312L9 13.125Q8.95 13 8.838 13t-.238.05l-.1.375l.95 1.7q-1.05.6-1.7 1.625T7 19Zm2.75-1.5q-.2 0-.35-.15T9.25 17q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15Zm4.5 0q-.2 0-.35-.15t-.15-.35q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15ZM4 22V2h10l6 6v14H4Zm9-13h5l-5-5v5Z"/>`),
		g.Group(children),
	)
}