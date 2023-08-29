package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoapBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSoapBubble0"><g fill="none" stroke-linecap="round" stroke-width="4"><ellipse cx="22" cy="30" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="16" ry="14"/><path stroke="#000" d="M26 24c1.333.167 4 1 5 5"/><circle cx="8" cy="8" r="4" fill="#fff" stroke="#fff" stroke-linejoin="round"/><circle cx="41" cy="13" r="3" fill="#fff" stroke="#fff" stroke-linejoin="round"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSoapBubble0)"/>`),
		g.Group(children),
	)
}