package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StereoPerspective(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 44V14L14 4h30v30L34 44H4Zm30-30v30M14 4v30M4 14h30M44 4L34 14M4 44l10-10m0 0h30"/>`),
		g.Group(children),
	)
}