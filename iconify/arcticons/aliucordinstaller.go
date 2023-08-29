package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aliucordinstaller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.426 18.003v.053a4.238 4.238 0 0 1-4.238 4.238h0a4.238 4.238 0 0 1-4.238-4.238v-4.318A4.238 4.238 0 0 1 29.188 9.5h0a4.238 4.238 0 0 1 4.238 4.238v.052m-18.852 8.466L18.82 9.5m4.07 12.794L18.82 9.5m2.709 8.514h-5.543m2.444 15.208l5.278 5.278l5.278-5.278M23.708 38.5V25.706"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}