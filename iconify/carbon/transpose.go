package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transpose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 26h-5v-2h5a5.005 5.005 0 0 0 5-5v-5h2v5a7.008 7.008 0 0 1-7 7zM8 30H4a2.002 2.002 0 0 1-2-2V14a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2v14a2.002 2.002 0 0 1-2 2zM4 14v14h4V14zm24-4H14a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h14a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM14 4v4h14V4z"/>`),
		g.Group(children),
	)
}