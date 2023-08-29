package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFileDate0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m30 4l10 10"/><circle cx="24" cy="28" r="8" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M23 25v4h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFileDate0)"/>`),
		g.Group(children),
	)
}