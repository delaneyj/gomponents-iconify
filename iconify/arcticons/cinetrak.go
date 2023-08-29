package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinetrak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.798 41.057a4.852 4.852 0 0 1-3.453 1.443H7.505A2.005 2.005 0 0 1 5.5 40.495V7.505A2.005 2.005 0 0 1 7.505 5.5h32.99A2.005 2.005 0 0 1 42.5 7.505v25.42a3.622 3.622 0 0 1-1.106 2.605Z"/><circle cx="12.947" cy="10.573" r="2.086" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.703 8.519l2.054 2.055l-2.054 2.054m7.184-4.109l2.054 2.055l-2.054 2.054m7.183-4.109l2.055 2.055l-2.055 2.054m-6.26 18.601a5.147 5.147 0 0 1-4.47 2.593h0a5.149 5.149 0 0 1-5.15-5.149v-3.346a5.149 5.149 0 0 1 5.149-5.15h0a5.147 5.147 0 0 1 4.466 2.585"/>`),
		g.Group(children),
	)
}