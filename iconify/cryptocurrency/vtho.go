package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vtho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm4.272-26.942h-5.763l-4.487 11.78h4.45l-3.94 10.104l12.18-14.3h-5.1l5.683-7.584h-3.023z"/>`),
		g.Group(children),
	)
}