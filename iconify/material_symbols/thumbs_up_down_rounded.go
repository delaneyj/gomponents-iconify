package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 14q-.825 0-1.413-.588T0 12V6.225q0-.4.15-.762t.425-.638L4.65.75Q4.9.5 5.275.462t.675.163q.3.2.45.537t.075.688L5.8 5h4.25q1.2 0 1.725.925T11.9 7.8l-2.125 4.975q-.25.575-.738.9T7.95 14H2Zm15.85 9.225q-.25-.2-.313-.5t-.012-.575L18.2 19H14q-1.2 0-1.763-.913T12.1 16.2l2.125-4.975q.25-.575.738-.9T16.05 10H22q.825 0 1.413.587T24 12v5.775q0 .4-.15.763t-.425.637L19.35 23.25q-.325.325-.75.288t-.75-.313Z"/>`),
		g.Group(children),
	)
}