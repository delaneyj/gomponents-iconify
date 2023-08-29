package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm10.994 44.508H21.006V17.492h21.279v5.139H26.932v6.16h14.094v5.039H26.932v7.461h16.063v5.217z"/>`),
		g.Group(children),
	)
}