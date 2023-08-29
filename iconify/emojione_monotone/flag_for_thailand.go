package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForThailand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-17.756 8.367h35.512a28.16 28.16 0 0 1 7.628 9.833H6.616a28.163 28.163 0 0 1 7.628-9.833m35.512 43.266H14.244A28.152 28.152 0 0 1 6.616 43.8h50.768a28.172 28.172 0 0 1-7.628 9.833"/>`),
		g.Group(children),
	)
}