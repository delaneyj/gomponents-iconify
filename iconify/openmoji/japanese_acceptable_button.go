package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseAcceptableButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36.5" cy="36.5" r="27.5" fill="#d0cfce"/><path d="M36.5 11A25.5 25.5 0 1 1 11 36.5A25.529 25.529 0 0 1 36.5 11m0-2A27.5 27.5 0 1 0 64 36.5A27.5 27.5 0 0 0 36.5 9Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 25h31M37.5 52.5h8v-26m-20 21v-14h12v10h-11"/>`),
		g.Group(children),
	)
}