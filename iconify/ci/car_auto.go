package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11h1.045m0 0h15.91m-15.91 0c.012-.052.025-.102.04-.153a4.04 4.04 0 0 1 .19-.467L5.822 6.9c.305-.687.458-1.032.7-1.284c.214-.223.476-.393.766-.498C7.617 5 7.993 5 8.746 5h6.508c.752 0 1.13 0 1.458.118c.29.105.552.275.765.498c.242.252.395.596.7 1.283l1.553 3.493c.099.223.15.337.185.455c.015.05.028.101.04.153m-15.91 0a2.021 2.021 0 0 0-.03.174C4 11.3 4 11.426 4 11.68V17m15.954-6H21m-1.046 0c.013.058.023.116.03.174c.016.124.016.25.016.5V17m0 0h-4m4 0v1a2 2 0 1 1-4 0v-1m0 0H8m0 0H4m4 0v1a2 2 0 1 1-4 0v-1"/>`),
		g.Group(children),
	)
}