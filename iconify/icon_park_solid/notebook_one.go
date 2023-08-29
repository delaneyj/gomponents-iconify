package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNotebookOne0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M8 6a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v36a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V6Z"/><path stroke="#000" stroke-linecap="round" d="M16 4v40m8-32h8m-8 8h8"/><path stroke="#fff" stroke-linecap="round" d="M10 4h12M10 44h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNotebookOne0)"/>`),
		g.Group(children),
	)
}