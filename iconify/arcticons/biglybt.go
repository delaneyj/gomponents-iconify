package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Biglybt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.631 15.363h3.448m-8.758 3.821H8.988a14.768 14.768 0 0 0 3.09 5.746H8.095c-2.112 0-3.508 1.51-2.126 2.892L21.33 42.207a4.788 4.788 0 0 0 6.658-.109L42.01 28.08c1.42-1.42.017-3.15-2.157-3.15h-3.325a14.768 14.768 0 0 0 3.09-5.747l-9.333.001m0-7.643h9.333c-.749-2.724-2.28-5.148-4.35-7.041h0h-21.93c-2.07 1.893-3.601 4.316-4.35 7.041H19.59"/><circle cx="34.797" cy="15.363" r=".75" fill="currentColor"/><circle cx="38.896" cy="15.349" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.285 20.136H21.49a2.946 2.946 0 0 1-2.946-2.946v-3.265a3.238 3.238 0 0 1 3.239-3.238h8.503v9.449Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.384 17.764h5.359c.186 0 .337-.15.337-.336V13.44a.336.336 0 0 0-.337-.336h-5.359a.704.704 0 0 0-.704.704v3.251c0 .39.316.704.704.704Z"/>`),
		g.Group(children),
	)
}