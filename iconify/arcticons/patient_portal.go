package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatientPortal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.5 19.606C26.689 21.413 18.372 29.732 12.224 42.93"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.677 4.662a43.47 43.47 0 0 0-1.942 15.783m2.54 10.056a46.456 46.456 0 0 0 8.885 12.837M18.663 25.644A30.89 30.89 0 0 0 3.5 22.993"/>`),
		g.Group(children),
	)
}