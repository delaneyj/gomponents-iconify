package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MerckPse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.424 8.151c-1.692 0-6.834 8.404-13.424 8.404S12.268 8.151 10.576 8.151H5.5v14.237a2.858 2.858 0 0 0 2.85 2.85h2.226V20.34c0-1.915.935-2.894 2.838-2.894c3.04 0 7.09 6.857 10.586 6.857s7.547-6.857 10.586-6.857c1.903 0 2.838.98 2.838 2.894v2.048a2.85 2.85 0 0 0 2.85 2.85H42.5V10.15a2 2 0 0 0-2-2h-3.075ZM14.862 39.849v-8h2.619c1.48 0 2.68 1.203 2.68 2.686s-1.2 2.687-2.68 2.687h-2.62m7.204 1.75c.49.639 1.105.877 1.961.877h1.185a1.996 1.996 0 0 0 1.995-1.996v-.009a1.996 1.996 0 0 0-1.995-1.995h-1.307a1.998 1.998 0 0 1-1.998-1.998h0c0-1.106.897-2.002 2.002-2.002h1.178c.856 0 1.471.237 1.962.876m2.09 3.124h2.608m1.392 4h-4v-8h4"/>`),
		g.Group(children),
	)
}