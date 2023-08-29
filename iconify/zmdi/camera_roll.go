package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRoll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 91h171v320H256q0 17-12.5 29.5T213 453H43q-18 0-30.5-12.5T0 411V91q0-18 12.5-30.5T43 48h21V27q0-9 6.5-15.5T85 5h86q8 0 14.5 6.5T192 27v21h21q18 0 30.5 12.5T256 91zm-43 277v-43h-42v43h42zm0-192v-43h-42v43h42zm86 192v-43h-43v43h43zm0-192v-43h-43v43h43zm85 192v-43h-43v43h43zm0-192v-43h-43v43h43z"/>`),
		g.Group(children),
	)
}