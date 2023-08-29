package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItFortyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><mask id="circleFlagsIt450" width="512" height="512" x="0" y="0" maskUnits="userSpaceOnUse"><circle cx="256" cy="256" r="256" fill="#fff"/></mask></defs><g mask="url(#circleFlagsIt450)"><path fill="#eee" d="M0 0h512v512H0V0Z"/><path fill="#496e2d" d="M136 128v240h240V128H136Zm16 224V247l203 105H152Zm208-161a166 166 0 0 1-104 0a166 166 0 0 0-104 0v-47h208v47Z"/><path fill="#d80027" d="M136 384h240v16H136z"/></g>`),
		g.Group(children),
	)
}