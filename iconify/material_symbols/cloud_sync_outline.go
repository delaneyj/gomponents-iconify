package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSyncOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-2h2.725q-1.275-1.1-2-2.65T4 12q0-2.8 1.7-4.938T10 4.25v2.1q-1.75.625-2.875 2.163T6 12q0 1.35.537 2.488T8 16.45V14h2v6H4Zm11 0q-1.25 0-2.125-.875T12 17q0-1.2.825-2.063t2.025-.912q.425-.9 1.263-1.463T18 12q1.325 0 2.288.863T21.45 15q1.05 0 1.8.725t.75 1.75q0 1.05-.725 1.788T21.5 20H15Zm2.9-9q-.175-1.025-.675-1.9T16 7.55V10h-2V4h6v2h-2.725q1.075.95 1.763 2.225T19.925 11H17.9ZM15 18h6.5q.2 0 .35-.15t.15-.35q0-.2-.15-.35T21.5 17h-1.75v-1.25q0-.725-.513-1.238T18 14q-.725 0-1.238.513t-.512 1.237V16H15q-.425 0-.713.288T14 17q0 .425.288.713T15 18Zm3-2Z"/>`),
		g.Group(children),
	)
}