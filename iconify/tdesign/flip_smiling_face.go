package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipSmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm6.67-3.5A4.998 4.998 0 0 1 12 6a4.998 4.998 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A2.995 2.995 0 0 0 12 8a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001l.5-.866ZM9 12v4H7v-4h2Zm8 0v4h-2v-4h2Z"/>`),
		g.Group(children),
	)
}