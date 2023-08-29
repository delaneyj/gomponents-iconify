package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOutdoor0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m4 42l14-32l10 24l4-12l12 20H4Zm33-28a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOutdoor0)"/>`),
		g.Group(children),
	)
}