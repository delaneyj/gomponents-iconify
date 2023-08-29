package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileFocus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileFocus0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke="#fff" d="m30 4l10 10"/><path fill="#000" stroke="#000" d="m24 19l2.523 5.527l6.037.692l-4.477 4.108l1.207 5.954L24 32.293l-5.29 2.988l1.207-5.954l-4.477-4.108l6.037-.692L24 19Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileFocus0)"/>`),
		g.Group(children),
	)
}