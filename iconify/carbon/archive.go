package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 19h4v2h-4z"/><path fill="currentColor" d="M6 2v26a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V2Zm18 26H8V16h16Zm0-14H8v-4h16ZM8 8V4h16v4Z"/>`),
		g.Group(children),
	)
}