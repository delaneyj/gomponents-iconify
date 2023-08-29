package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sporting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSporting0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="8" r="4" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M7 18h12v16m22-16H29v26"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSporting0)"/>`),
		g.Group(children),
	)
}