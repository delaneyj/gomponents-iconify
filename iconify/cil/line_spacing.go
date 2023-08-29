package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 48h256v32H240zm0 128h256v32H240zm0 128h256v32H240zm0 128h256v32H240zM24 368v30.19l88.937 97.728L200 398.089V368h-72.8V144H200v-30.19l-88.937-97.728L24 113.911V144h71.2v224Zm44.538-256l42.791-48.082L155.086 112Zm86.924 288l-42.791 48.082L68.914 400Z"/>`),
		g.Group(children),
	)
}