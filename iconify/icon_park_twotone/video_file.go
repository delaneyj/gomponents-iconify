package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTVideoFile0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path d="m30 4l10 10"/><path fill="#555" d="M14 21h13.493v3.5L34 22v11l-6.507-2.5V34H14V21Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTVideoFile0)"/>`),
		g.Group(children),
	)
}