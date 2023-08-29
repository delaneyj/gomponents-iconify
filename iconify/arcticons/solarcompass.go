package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Solarcompass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.05 34.423a11.995 11.995 0 1 0-4.474-4.473"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 24l-10.024 5.179l4.845 4.845L24 24zM8.797 39.203l10.024-5.179l-4.845-4.845l-5.179 10.024zM24 9.482V2.5m0 43v-6.981m-7.259-27.092L13.25 5.38m21.5 37.24l-3.491-6.047M11.427 16.741L5.38 13.25m37.24 21.5l-6.047-3.491M9.482 24H2.5m43 0h-6.981m4.101-10.75l-6.047 3.491M34.75 5.38l-3.491 6.047M38.519 24H45.5m-43 0h6.982"/>`),
		g.Group(children),
	)
}