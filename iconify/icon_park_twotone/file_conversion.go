package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileConversion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFileConversion0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path d="m30 4l10 10M17 25h14m-14 6h14m0-6l-5-5m-4 16l-5-5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFileConversion0)"/>`),
		g.Group(children),
	)
}