package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateClockwiseFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 30H16a2.002 2.002 0 0 1-2-2V16a2.002 2.002 0 0 1 2-2h12a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2zM15 2l-1.41 1.41L16.17 6H11a7.008 7.008 0 0 0-7 7v5h2v-5a5.006 5.006 0 0 1 5-5h5.17l-2.58 2.59L15 12l5-5z"/>`),
		g.Group(children),
	)
}