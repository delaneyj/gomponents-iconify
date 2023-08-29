package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 99.87v1000.26h100.26V99.87H0zm186.135 0v1000.26h24.089V99.87h-24.089zm57.883 0v1000.26h76.219V99.87h-76.219zm153.658 0v1000.26h32.525V99.87h-32.525zm118.449 0v1000.26h24.041V99.87h-24.041zm34.379 0v1000.26h99.626V99.87h-99.626zm159.411 0v1000.26h50.18V99.87h-50.18zm149.025 0v1000.26h11.118V99.87H858.94zm87.63 0v1000.26h33.501V99.87H946.57zm119.376 0v1000.26h24.09V99.87h-24.09zm34.428 0v1000.26H1200V99.87h-99.626z"/>`),
		g.Group(children),
	)
}