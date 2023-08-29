package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryProfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V2h4v2h2q.425 0 .713.288T17 5v5q-2.875 0-4.925 2.025t-2.05 4.95q0 1.425.537 2.725t1.563 2.3H8Zm8 0l-.3-1.5q-.3-.125-.563-.275t-.512-.35l-1.45.5l-1-1.725l1.125-1q-.05-.375-.05-.638t.05-.637l-1.125-1l1-1.725l1.45.5q.25-.2.513-.363t.562-.287L16 12h2l.3 1.5q.3.125.563.275t.512.35l1.45-.5l1 1.725l-1.125 1q.05.375.05.638t-.05.637l1.125 1l-1 1.725l-1.45-.475q-.25.2-.513.35t-.562.275L18 22h-2Zm1-3q.825 0 1.413-.588T19 17q0-.825-.588-1.413T17 15q-.825 0-1.413.588T15 17q0 .825.588 1.413T17 19Z"/>`),
		g.Group(children),
	)
}