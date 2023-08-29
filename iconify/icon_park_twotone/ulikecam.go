package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ulikecam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUlikecam0"><path fill="#555" stroke="#fff" stroke-width="4" d="M22 44c9.941 0 18-8.059 18-18S31.941 8 22 8S4 16.059 4 26s8.059 18 18 18Zm19-34a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUlikecam0)"/>`),
		g.Group(children),
	)
}