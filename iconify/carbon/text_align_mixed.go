package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignMixed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 4H9v2h5v2h-4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6V6a2.002 2.002 0 0 0-2-2zm0 8h-4v-2h4zm8 14v-8h-2v1h-2v2h2v5h-2v2h6v-2h-2zM2 2h2v14H2zm26 14h2v14h-2z"/>`),
		g.Group(children),
	)
}