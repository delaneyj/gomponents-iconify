package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOutdoor0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m4 42l14-32l10 24l4-12l12 20H4Zm33-28a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOutdoor0)"/>`),
		g.Group(children),
	)
}