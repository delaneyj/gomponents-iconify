package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPlayOne0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M15 24V11.876l10.5 6.062L36 24l-10.5 6.062L15 36.124V24Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPlayOne0)"/>`),
		g.Group(children),
	)
}