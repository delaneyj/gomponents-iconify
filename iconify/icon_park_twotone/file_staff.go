package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileStaff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFileStaff0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path d="m30 4l10 10"/><circle cx="24" cy="24" r="4" fill="#555"/><path d="M32 36a8 8 0 1 0-16 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFileStaff0)"/>`),
		g.Group(children),
	)
}