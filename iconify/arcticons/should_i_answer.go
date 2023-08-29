package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShouldIAnswer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.049 28.767a28.677 28.677 0 0 1-6.376-8.516c1.55-1.59 3-3.109 3.589-3.638a3.109 3.109 0 0 0 .58-3a61.816 61.816 0 0 1-1.11-6.307a2 2 0 0 0-2.194-1.783q-.053.005-.105.014H6.916a1.5 1.5 0 0 0-1.36 1.369a28.852 28.852 0 0 0 1.852 11.265M23.72 37.935h0c2 1.06 9.947 5.058 17.374 4.508a1.5 1.5 0 0 0 1.39-1.359v-8.507a2 2 0 0 0-1.665-2.286a2.086 2.086 0 0 0-.105-.013a61.99 61.99 0 0 1-6.298-1.11a3.109 3.109 0 0 0-2.999.58c-.53.54-2.059 2.049-3.648 3.608q-.805-.395-1.58-.838M8.567 21.115a40.4 40.4 0 0 0 1.598 3.285h0v.06l.12.23h0a35.427 35.427 0 0 0 6.625 8.302"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.49 25.04c-1.534.357-2.664.387-3.106-.227c-1.739-2.41 3.644-4.654 2.467-6.897c-.774-1.474-9.309 3.309-6.617 8.411c2.18 4.134 6.815 4.79 11.244 4.182"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.049 28.767c-2.114 3-9.292 7.298-9.404 13.728h12.897c-4.766-4.392 7.066-14.43 3.645-18.495c-2.147-2.552-7.083-2.41-11.847-.807"/>`),
		g.Group(children),
	)
}