package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileSearchTwo0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m30 4l10 10"/><circle cx="22" cy="26" r="6" fill="#000" stroke="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m27 30l5 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileSearchTwo0)"/>`),
		g.Group(children),
	)
}