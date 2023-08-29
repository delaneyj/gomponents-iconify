package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.516 7.167H3.482L16 14.275l12.516-7.108zM16.74 17.303a1.494 1.494 0 0 1-1.48 0L2.5 10.06v14.773h27V10.06l-12.76 7.243z"/>`),
		g.Group(children),
	)
}