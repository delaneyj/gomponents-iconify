package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CutOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.05 13.733l-1-1.733l-10.122 5.846l-.997-.576a2.952 2.952 0 0 0 .667-.769A3 3 0 1 0 3.5 17.599L5.928 19L3.5 20.402a3.034 3.034 0 1 0 3.44.323l.988-.57L14.59 24l1-1.73L9.928 19zM4.034 15.26a1 1 0 1 1 .466.607a.993.993 0 0 1-.466-.607zM5 22a1 1 0 1 1-.865 1.5A1 1 0 0 1 5 22zm12 4h4v2h-4zm-7 0h4v2h-4z"/><path fill="currentColor" d="M28 28h-4v-2h4V4H7v4H5V4a2.002 2.002 0 0 1 2-2h21a2.002 2.002 0 0 1 2 2v22a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}