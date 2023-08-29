package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 16q-1.875 0-3.188-1.313T2 11.5q0-1.875 1.313-3.188T6.5 7q1.875 0 3.188 1.313T11 11.5q0 .675-.2 1.3t-.55 1.2h3.5q-.35-.575-.55-1.2t-.2-1.3q0-1.875 1.313-3.188T17.5 7q1.875 0 3.188 1.313T22 11.5q0 1.875-1.313 3.188T17.5 16h-11Zm0-2q1.05 0 1.775-.725T9 11.5q0-1.05-.725-1.775T6.5 9q-1.05 0-1.775.725T4 11.5q0 1.05.725 1.775T6.5 14Zm11 0q1.05 0 1.775-.725T20 11.5q0-1.05-.725-1.775T17.5 9q-1.05 0-1.775.725T15 11.5q0 1.05.725 1.775T17.5 14Z"/>`),
		g.Group(children),
	)
}