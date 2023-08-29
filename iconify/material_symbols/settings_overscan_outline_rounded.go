package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsOverscanOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 10.6q0-.175-.15-.237t-.275.062L5.35 11.65q-.15.15-.15.35t.15.35l1.225 1.225q.125.125.275.063T7 13.4v-2.8Zm6.4 4.4h-2.8q-.175 0-.237.15t.062.275l1.225 1.225q.15.15.35.15t.35-.15l1.225-1.225q.125-.125.063-.275T13.4 15Zm0-6q.175 0 .238-.15t-.063-.275L12.35 7.35Q12.2 7.2 12 7.2t-.35.15l-1.225 1.225q-.125.125-.062.275T10.6 9h2.8Zm3.6 1.6v2.8q0 .175.15.238t.275-.063l1.225-1.225q.15-.15.15-.35t-.15-.35l-1.225-1.225q-.125-.125-.275-.062T17 10.6ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}