package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolicyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14q.825 0 1.413-.588T14 12q0-.825-.588-1.413T12 10q-.825 0-1.413.588T10 12q0 .825.588 1.413T12 14Zm0 7.85q-.275 0-.513-.038t-.437-.112q-3.1-1.15-5.075-4.1T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 1.5-.4 3.025T18.4 17l-2.95-2.95q.275-.475.413-.988T16 12q0-1.65-1.175-2.825T12 8q-1.65 0-2.825 1.175T8 12q0 1.65 1.175 2.825T12 16q.525 0 1.038-.138T14 15.45l3.225 3.225q-.875 1.05-1.95 1.813T12.95 21.7q-.2.075-.437.113T12 21.85Z"/>`),
		g.Group(children),
	)
}