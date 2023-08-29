package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Planisphere(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="12.883" cy="23.018" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.629" cy="24.642" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.056" cy="24.642" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.633" cy="31.825" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="6.25" cy="29.475" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="41.75" cy="26.875" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="40.7" cy="16.175" r="1.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.356 24.915a29.1 29.1 0 0 0 3.978.03M14.55 23.54a43.557 43.557 0 0 0 4.339.955M7.48 28.231l4.124-4.02m17.011 2.088l1.418 3.882m10.838-12.265l.707 7.209m-9.347 5.989l7.922-3.528m-.907-10.437l-9.736 6.519"/>`),
		g.Group(children),
	)
}