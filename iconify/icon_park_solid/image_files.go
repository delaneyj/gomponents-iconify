package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSImageFiles0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke="#fff" d="m30 4l10 10"/><circle cx="18" cy="17" r="4" fill="#000" stroke="#000"/><path fill="#000" stroke="#000" d="M15 28v9h18V21l-9.51 10.5L15 28Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSImageFiles0)"/>`),
		g.Group(children),
	)
}