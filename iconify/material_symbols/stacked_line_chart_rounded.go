package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedLineChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 15.5l-5.25 5.25q-.325.325-.75.325t-.75-.325q-.325-.325-.325-.75t.325-.75L8.8 13.2q.3-.3.7-.3t.7.3l3.3 3.3l6.4-7.225q.275-.325.713-.325t.737.3q.275.275.287.663t-.262.687L14.2 18.7q-.275.325-.712.338t-.738-.288L9.5 15.5Zm0-6l-5.25 5.25q-.325.325-.75.325t-.75-.325q-.325-.325-.325-.75t.325-.75L8.8 7.2q.3-.3.7-.3t.7.3l3.3 3.3l6.4-7.225q.275-.325.713-.325t.737.3q.275.275.287.662t-.262.688L14.2 12.7q-.275.325-.712.338t-.738-.288L9.5 9.5Z"/>`),
		g.Group(children),
	)
}