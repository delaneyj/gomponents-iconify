package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 432v-43h43v43h-43zm85-256v-43h43v43h-43zM0 91q0-18 12.5-30.5T43 48h85v43H43v298h85v43H43q-18 0-30.5-12.5T0 389V91zm341-43q18 0 30.5 12.5T384 91h-43V48zM171 475V5h42v470h-42zm170-128v-43h43v43h-43zM256 91V48h43v43h-43zm85 170v-42h43v42h-43zm0 171v-43h43q0 18-12.5 30.5T341 432z"/>`),
		g.Group(children),
	)
}