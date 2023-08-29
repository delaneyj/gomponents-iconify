package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M249 42Q142 41 59 97l43 53q49-29 107-38q108-15 187 38l42-53q-83-56-189-55zm-.5-39Q387 3 497 88L249 397L0 88Q110 3 248.5 3z"/>`),
		g.Group(children),
	)
}