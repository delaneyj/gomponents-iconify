package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileSettings0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m30 4l10 10"/><circle cx="24" cy="27" r="5" fill="#000" stroke="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M24 19v3m0 10v3m5.828-14l-2.121 2.121M19.828 31l-2.121 2.121M18 21l2.121 2.121M28 31l2.121 2.121M16 27h3m10 0h3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileSettings0)"/>`),
		g.Group(children),
	)
}