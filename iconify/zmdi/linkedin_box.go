package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M363 365V244q0-31-22-53t-53-22q-15 0-30 8.5T235 199v-26h-64v192h64V252q0-13 9-22.5t22.5-9.5t23 9.5T299 252v113h64zM96 137q16 0 27.5-11T135 99t-11.5-27.5T96 60T68.5 71.5T57 99t11.5 27T96 137zm32 228V173H64v192h64zM384 3q18 0 30.5 12.5T427 45v342q0 17-12.5 29.5T384 429H43q-18 0-30.5-12.5T0 387V45q0-17 12.5-29.5T43 3h341z"/>`),
		g.Group(children),
	)
}