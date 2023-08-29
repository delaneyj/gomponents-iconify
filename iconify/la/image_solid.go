package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 5v22h28V5zm2 2h24v13.906l-5.281-5.312l-.719-.719l-4.531 4.531l-5.75-5.812l-.719-.719l-7 7zm20 2a1.999 1.999 0 1 0 0 4a1.999 1.999 0 1 0 0-4zm-13 6.719L20.188 25H4v-2.281zm11 2l6 6V25h-4.969l-4.156-4.188z"/>`),
		g.Group(children),
	)
}