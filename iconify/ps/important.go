package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Important(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M43 280h42q18 0 30.5-12.5T128 237V45q0-17-12.5-29.5T85 3H43Q25 3 12.5 15.5T0 45v192q0 18 12.5 30.5T43 280zm0-235h42v192H43V45zM0 387q0 17 12.5 29.5T43 429h42q18 0 30.5-12.5T128 387v-22q0-17-12.5-29.5T85 323H43q-18 0-30.5 12.5T0 365v22zm43-22h42v22H43v-22z"/>`),
		g.Group(children),
	)
}