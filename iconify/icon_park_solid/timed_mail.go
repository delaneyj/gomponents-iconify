package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimedMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTimedMail0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M44 35V9H4v28h22"/><circle cx="35" cy="35" r="9" fill="#fff" stroke="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M34 32v4h4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m4 9l20 13L44 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTimedMail0)"/>`),
		g.Group(children),
	)
}