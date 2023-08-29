package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bvgfahrinfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="M8.978 25.258v-14.48m5.512 14.48v-14.48M20 25.258V10.807m7.997 14.451V10.796m5.511 14.462V10.784m5.512 14.474v-14.48m-32.615-.002h35.19m-19.284.004V5.5m3.492 5.28V5.5m-6.984 5.312v-4.4m10.478 4.351v-4.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.234 36.304a2.795 2.795 0 0 1 0 5.59h-4.57v-11.18h4.57a2.795 2.795 0 0 1 0 5.59Zm12.353-5.684l-3.601 11.181l-3.74-11.181m16.877 3.774a3.77 3.77 0 0 0-3.878-3.774a3.924 3.924 0 0 0-3.462 3.913v3.494a3.741 3.741 0 0 0 3.74 3.774h0a3.741 3.741 0 0 0 3.739-3.774h-3.74"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" d="M15.234 36.304h-4.57"/>`),
		g.Group(children),
	)
}