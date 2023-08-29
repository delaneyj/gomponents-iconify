package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnpublishedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-1.55.425-2.938T3.65 6.476L1.375 4.2q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3l-2.25-2.225q-1.2.8-2.587 1.225T12 22Zm8.375-4.5l-5.325-5.35l1.9-1.9q.275-.275.275-.7t-.275-.7q-.3-.3-.713-.3t-.687.3l-1.9 1.925l-7.15-7.15q1.2-.775 2.588-1.2T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 1.525-.425 2.913t-1.2 2.587ZM7.05 13.05L9.9 15.9q.3.3.7.3t.7-.3l.9-.9l-1.4-1.4l-.2.2l-2.15-2.15q-.275-.275-.7-.275t-.7.275q-.275.275-.275.7t.275.7Z"/>`),
		g.Group(children),
	)
}