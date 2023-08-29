package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V20q0 .825-.588 1.413T18 22H6Zm7-14q0 .425.288.713T14 9h4l-5-5v4Zm-2.075 7.125l-1.4-1.4q-.15-.15-.325-.225t-.362-.075q-.188 0-.375.075t-.338.225q-.3.3-.3.713t.3.712l2.125 2.15q.15.15.325.213t.375.062q.2 0 .375-.063t.325-.212l4.225-4.225q.3-.3.3-.725t-.3-.725q-.3-.3-.725-.3t-.725.3l-3.5 3.5Z"/>`),
		g.Group(children),
	)
}