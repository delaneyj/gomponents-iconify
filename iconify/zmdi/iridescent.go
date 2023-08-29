package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iridescent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M31 301V173h299v128H31zM159 4h43v63h-43V4zm171 53l31 30l-39 38l-30-30zM202 471h-43v-63h43v63zm159-83l-31 30l-38-39l30-30zM0 87l30-30l38 38l-30 30zm30 331L0 387l38-38l30 30z"/>`),
		g.Group(children),
	)
}