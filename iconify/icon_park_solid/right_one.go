package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRightOne0"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m20 12l12 12l-12 12V12Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRightOne0)"/>`),
		g.Group(children),
	)
}