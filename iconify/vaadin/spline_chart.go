package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplineChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M1 15V0H0v16h16v-1H1z"/><path fill="currentColor" d="M12 5c-.69 1-1.41 2-2 2l-.087.001c-.601 0-1.164-.16-1.65-.44a4.519 4.519 0 0 0-2.2-.562h-.067a5.83 5.83 0 0 0-3.991 1.993l-.006 2.347c.77-1.12 2.32-2.84 4-2.84h.048c.579 0 1.121.156 1.587.428a4.682 4.682 0 0 0 2.264.573l.106-.001c1.395 0 2.335-1.32 3.245-2.6s1.75-2.4 2.75-2.4v-1.5c-1.81 0-3 1.61-4 3z"/>`),
		g.Group(children),
	)
}