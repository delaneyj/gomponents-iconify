package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FZeroKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFZeroKey0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><rect width="8" height="16" x="26" y="16" fill="#fff" stroke="#000" rx="4"/><path stroke="#000" d="M21 16h-7v16m0-8h7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFZeroKey0)"/>`),
		g.Group(children),
	)
}