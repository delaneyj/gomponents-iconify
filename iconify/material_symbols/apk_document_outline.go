package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApkDocumentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19h10q-.1-1.225-.75-2.25t-1.7-1.625l.95-1.7q.05-.1.025-.225t-.15-.175q-.1-.05-.213-.025t-.162.125l-.975 1.75q-.5-.2-1-.313T12 14.45q-.525 0-1.025.113t-1 .312L9 13.125Q8.95 13 8.838 13t-.238.05l-.1.375l.95 1.7q-1.05.6-1.7 1.625T7 19Zm2.75-1.5q-.2 0-.35-.15T9.25 17q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15Zm4.5 0q-.2 0-.35-.15t-.15-.35q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15ZM6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8l6 6v12q0 .825-.588 1.413T18 22H6Zm7-13V4H6v16h12V9h-5ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}