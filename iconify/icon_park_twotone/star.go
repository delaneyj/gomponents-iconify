package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTStar0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m23.999 5l-6.113 12.478L4 19.49l10.059 9.834L11.654 43L24 36.42L36.345 43L33.96 29.325L44 19.491l-13.809-2.013L24 5Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTStar0)"/>`),
		g.Group(children),
	)
}