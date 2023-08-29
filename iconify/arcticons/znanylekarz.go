package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Znanylekarz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41 20.812l-1.669-.042a1.778 1.778 0 0 1-1.764-1.766v-1.845m-3.547 3.73l1.669.04a1.778 1.778 0 0 1 1.766 1.767v1.845M7 21.85v4.832h2.416m16.707.08h2.417m-2.416-4.833h2.416m-2.416 2.416h1.57m-1.57-2.416v4.832m-6.33-.072v-4.833l2.416 4.832l2.416-4.832v4.832m-9.433-4.8l3.201 4.832m0-4.832l-3.201 4.832m14.79-.057V21.83h.786a2.424 2.424 0 0 1 2.416 2.416h0a2.424 2.424 0 0 1-2.416 2.416h-.785m-19.275-4.774v3.201a1.572 1.572 0 1 0 3.142 0v-3.201"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 35.5v-23a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v23a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}