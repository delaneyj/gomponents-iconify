package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SevenMWorkout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="38.041" cy="13.692" r="3.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="41.508" cy="28.885" r="3.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.792" cy="41.069" r="3.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.208" cy="41.069" r="3.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="6.492" cy="28.885" r="3.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.959" cy="13.692" r="3.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="6.93" r="3.991" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}