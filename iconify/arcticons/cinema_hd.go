package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CinemaHd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 20.8l-4.878-10.066a3.942 3.942 0 0 1-5.26-1.832L6.341 21.96A3.932 3.932 0 0 1 4.5 27.21l4.878 10.065a3.934 3.934 0 0 1 5.26 1.823L41.659 26.04a3.932 3.932 0 0 1 1.841-5.25l-4.878-10.065M13.573 25.404l2.992 5.985m7.378-3.582l3.083-1.451m-5.984-4.534l2.992-1.451m-1.542 4.443l1.995-.906m-3.445-2.086l2.901 5.985"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.401 25.576l-2.916-6.043l5.937 4.585l.105-7.501l2.917 6.043m-16.299 8.067l-2.916-6.042l7.448 3.855l-2.916-6.043m20.155-2.437l-4.88-5.094l.877 7.026m-.302-2.462l2.624-1.225M14.12 30.048a2.261 2.261 0 1 1-4.07 1.968l-.989-2.04a2.25 2.25 0 0 1 1.052-3.02a2.184 2.184 0 0 1 2.947 1.088"/>`),
		g.Group(children),
	)
}