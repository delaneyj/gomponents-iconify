package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusSimulatorIndonesia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5a28.494 28.494 0 0 1 4.961.416c1.938.326 3.825.905 5.613 1.72a5.48 5.48 0 0 1 2.533 1.89c.327.62.53 1.297.596 1.994a89.092 89.092 0 0 1 1.417 10.168V37.64h-2.542v2.542a2.174 2.174 0 1 1-4.337 0V37.64H15.76v2.542a2.174 2.174 0 1 1-4.338 0V37.64H8.88V21.688a91.25 91.25 0 0 1 1.39-10.168a5.33 5.33 0 0 1 .595-1.994a5.48 5.48 0 0 1 2.532-1.89a22.556 22.556 0 0 1 5.613-1.72A28.505 28.505 0 0 1 24 5.501ZM13.605 28.473a2.278 2.278 0 1 0 2.267 2.288s0 0 0 0v-.02a2.267 2.267 0 0 0-2.267-2.268h0Zm20.79 0a2.278 2.278 0 1 0 2.277 2.278v-.01a2.259 2.259 0 0 0-2.249-2.267h-.028ZM11.422 37.64h25.156M8.88 21.688h30.24M18.988 32.406h10.024m-10.024-3.31h10.024"/>`),
		g.Group(children),
	)
}