package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApkInstall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 22l-4-4l1.4-1.425L19 18.15V14h2v4.15l1.6-1.575L24 18l-4 4ZM4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h8l6 6v4.25h-3V22H4Zm7-13h5l-5-5v5ZM5 19h10q-.1-1.225-.75-2.25t-1.7-1.625l.95-1.7q.05-.1.025-.225t-.15-.175q-.1-.05-.213-.025t-.162.125l-.975 1.75q-.5-.2-1-.313T10 14.45q-.525 0-1.025.113t-1 .312L7 13.125Q6.95 13 6.837 13t-.237.05l-.1.375l.95 1.7q-1.05.6-1.7 1.625T5 19Zm2.75-1.5q-.2 0-.35-.15T7.25 17q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15Zm4.5 0q-.2 0-.35-.15t-.15-.35q0-.2.15-.35t.35-.15q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15Z"/>`),
		g.Group(children),
	)
}