package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDisplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFileDisplay0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m30 4l10 10"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 34c4.97 0 9-6 9-6s-4.03-6-9-6s-9 6-9 6s4.03 6 9 6Z"/><path fill="#fff" d="M24 30a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFileDisplay0)"/>`),
		g.Group(children),
	)
}