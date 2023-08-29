package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApkInstallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h8l6 6v4.25h-2V9h-5V4H4v16h11v2H4Zm0-2V4v16Zm1-1q.1-1.225.75-2.25t1.7-1.625l-.95-1.7q0-.025.1-.375q.125-.05.238-.05t.162.125l.975 1.75q.5-.2 1-.313T10 14.45q.525 0 1.025.113t1 .312l.975-1.75l.375-.1q.125.05.15.175t-.025.225l-.95 1.7q1.05.6 1.7 1.625T15 19H5Zm2.75-1.5q.2 0 .35-.15t.15-.35q0-.2-.15-.35t-.35-.15q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Zm4.5 0q.2 0 .35-.15t.15-.35q0-.2-.15-.35t-.35-.15q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15ZM20 22l-4-4l1.4-1.425L19 18.15V14h2v4.15l1.6-1.575L24 18l-4 4Z"/>`),
		g.Group(children),
	)
}