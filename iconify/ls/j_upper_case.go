package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M317 529V0h72v527c0 66-28 125-72 166c-41 38-96 61-156 61c-62 0-120-25-161-66l52-51c28 29 66 46 109 46c85 0 155-69 156-154z"/>`),
		g.Group(children),
	)
}