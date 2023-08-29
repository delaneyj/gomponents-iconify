package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusTheTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMinusTheTop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M5 41V21a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2Z"/><path d="M43 7v20a2 2 0 0 1-2 2H31a2 2 0 0 1-2-2v-6a2 2 0 0 0-2-2h-6a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMinusTheTop0)"/>`),
		g.Group(children),
	)
}