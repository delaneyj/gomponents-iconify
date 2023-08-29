package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m221.7 82.83l-176.1-54A20 20 0 0 0 20 48v144a20 20 0 0 0 19.94 20a20.38 20.38 0 0 0 5.66-.81l78.4-24.05V192a20 20 0 0 0 20 20h32a20 20 0 0 0 20-20v-26.95l25.7-7.88A20.12 20.12 0 0 0 236 138v-36a20.12 20.12 0 0 0-14.3-19.17ZM124 162l-80 24.58V53.42L124 78Zm48 26h-24v-8.22l24-7.36Zm40-53l-64 19.63v-69.3L212 105Z"/>`),
		g.Group(children),
	)
}