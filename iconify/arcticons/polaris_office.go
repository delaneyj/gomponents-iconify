package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolarisOffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.136 7.89L43.5 18.739L35.472 40.11H14.914l9.222-32.22zm-9.222 32.22L4.5 18.955l6.606 1.836"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.914 40.11L9.707 13.694l6.238 2.326"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.703 12.897L16.216 9.68l-1.302 30.43m10.849-25.385l14.32 6.671M24.57 18.784l14.337 5.74m-15.531-1.681l14.356 4.809m-15.549-.75l14.374 3.878"/>`),
		g.Group(children),
	)
}