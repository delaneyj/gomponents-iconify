package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.431 2 2 15.432 2 32s13.432 30 30 30c16.568 0 30-13.432 30-30S48.568 2 32 2zm15 47h-6V24l-9 9l-9-9v25h-6V15h6l9 9l9-9h6v34z"/>`),
		g.Group(children),
	)
}