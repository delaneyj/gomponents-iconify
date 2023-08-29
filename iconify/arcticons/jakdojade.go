package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jakdojade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="9.04" cy="24.35" r="4.54" fill="none" stroke="currentColor" stroke-miterlimit="10"/><circle cx="37.6" cy="12.39" r="5.9" fill="none" stroke="currentColor" stroke-miterlimit="10"/><circle cx="19.19" cy="37.99" r="3.52" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M16.29 36a24.4 24.4 0 0 1-5.6-7.41m22.57-12.22A77.83 77.83 0 0 0 20.58 34.7"/><circle cx="9.04" cy="24.35" r=".75" fill="currentColor"/><circle cx="19.19" cy="37.99" r=".75" fill="currentColor"/><circle cx="37.6" cy="12.39" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}