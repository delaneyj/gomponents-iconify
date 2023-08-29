package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0zm2 411l114 115c14 14 37 14 51 0c13-13 13-36 0-50L411 361l114-115c13-14 13-37 0-50c-14-14-37-14-51 0L360 311L246 196c-14-14-37-14-51 0c-13 13-13 36 0 50l114 115l-114 115c-13 14-13 37 0 50c14 14 37 14 51 0z"/>`),
		g.Group(children),
	)
}