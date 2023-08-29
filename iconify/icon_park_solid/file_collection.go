package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCollection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileCollection0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2Z"/><path stroke="#fff" d="m30 4l10 10"/><path fill="#000" stroke="#000" d="M20.85 21C18.724 21 17 23.009 17 25.486c0 4.487 4.55 8.565 7 9.514c2.45-.949 7-5.027 7-9.514C31 23.01 29.276 21 27.15 21c-1.302 0-2.453.753-3.15 1.906C23.303 21.753 22.152 21 20.85 21Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileCollection0)"/>`),
		g.Group(children),
	)
}