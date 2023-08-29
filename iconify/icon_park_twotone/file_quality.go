package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileQuality(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFileQuality0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path d="m30 4l10 10"/><path fill="#555" d="M19.2 21h9.6l3.2 4.118L24 35l-8-9.882L19.2 21Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFileQuality0)"/>`),
		g.Group(children),
	)
}