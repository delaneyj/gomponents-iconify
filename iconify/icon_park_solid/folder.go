package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFolder0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M7 6a2 2 0 0 1 2-2h30a2 2 0 0 1 2 2v36a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V6Z"/><path stroke="#000" d="M16 29h4m-4 6h10M8 5s3.765 13 16 13S40 5 40 5"/><circle cx="24" cy="18" r="4" fill="#000" stroke="#000"/><path stroke="#fff" d="M15 4H9a2 2 0 0 0-2 2v6m26-8h6a2 2 0 0 1 2 2v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFolder0)"/>`),
		g.Group(children),
	)
}