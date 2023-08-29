package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalWellbeing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.105 40.226A21.5 21.5 0 1 1 20.637 2.765m5.983-.105a21.5 21.5 0 0 1 15.204 33.363"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.248c-3.464-10.938-21.931 2.46 0 15.534m0-15.534c3.464-10.938 21.931 2.46 0 15.534"/>`),
		g.Group(children),
	)
}