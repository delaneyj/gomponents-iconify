package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVoice0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="14" height="27" x="17" y="4" fill="#555" rx="7"/><path stroke-linecap="round" d="M9 23c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15M24 38v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVoice0)"/>`),
		g.Group(children),
	)
}