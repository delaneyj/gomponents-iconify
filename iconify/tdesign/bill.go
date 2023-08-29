package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v2h-2v18.08l-1.625-1.3l-1.841-1.472l-1.62 1.44l-.664.59l-.664-.59L12 19.337l-1.586 1.41l-.664.59l-.664-.59l-1.62-1.44l-1.841 1.473L4 22.08V4H2V2Zm4 2v13.92l.875-.7l.659-.528l.63.56l1.586 1.41l1.586-1.41l.664-.59l.664.59l1.586 1.41l1.586-1.41l.63-.56l.659.527l.875.7V4H6Zm2 3h8v2H8V7Zm2 4H9v2h6v-2h-5Z"/>`),
		g.Group(children),
	)
}