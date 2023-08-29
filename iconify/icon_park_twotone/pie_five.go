package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPieFive0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#555" d="M24 4a20 20 0 1 1-13.69 34.58L24 24V4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPieFive0)"/>`),
		g.Group(children),
	)
}