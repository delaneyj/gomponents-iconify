package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nurse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.2 12a4.076 4.076 0 0 0-3.931-2.37L8 13.53l-3.28-3.9a4.16 4.16 0 0 0-3.909 2.345a9.072 9.072 0 0 0-.808 2.988L2 15v1h12v-1h2a9.199 9.199 0 0 0-.824-3.057z"/><path fill="currentColor" d="M6.57 8.73a.82.82 0 0 1-.685.729L8 12l2.12-2.52a.823.823 0 0 1-.69-.727c0-.613.8-.413 1.43-1.453C10.86 7.27 13.74 0 8 0S5.14 7.27 5.14 7.27c.63 1.05 1.44.85 1.43 1.46z"/>`),
		g.Group(children),
	)
}