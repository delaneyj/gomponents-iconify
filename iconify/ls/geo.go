package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 161 359 359S556 717 358 717S0 557 0 359S160 0 358 0zm25 511v115c129-11 232-115 243-243H512v-47h114C616 206 512 103 383 92v115h-47V92C206 102 101 206 91 336h116v47H91c11 129 115 233 245 244V511h47z"/>`),
		g.Group(children),
	)
}