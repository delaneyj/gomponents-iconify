package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDisplayOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFileDisplayOne0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 23v-9L31 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h12"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M33 41c4.97 0 9-6 9-6s-4.03-6-9-6s-9 6-9 6s4.03 6 9 6Z"/><path fill="#fff" d="M33 37a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 4v10h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFileDisplayOne0)"/>`),
		g.Group(children),
	)
}