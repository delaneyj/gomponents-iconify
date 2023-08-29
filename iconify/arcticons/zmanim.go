package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zmanim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.67 6.147h2.66M41.853 22.67v2.66M25.33 41.853h-2.66M6.147 25.33v-2.66M25.329 2.505v3.642m0 35.706v3.593m-2.658 0v-3.593m0-35.706V2.554M2.554 22.671h3.593m35.706 0h3.642m-.049 2.658h-3.593m-35.706 0H2.554m42.946.004C44.828 36.18 36.18 44.828 25.333 45.5m-2.666 0C11.82 44.828 3.172 36.18 2.5 25.333m0-2.666C3.172 11.82 11.82 3.172 22.667 2.5m2.666 0c10.847.672 19.495 9.32 20.167 20.167m-24.79 4.225l18.602-11.936m-25.08 4.113l12.676 7.318m1.78.587c.54-2.842 3.001-4.896 5.866-4.896c3.07 0 5.641 2.352 5.944 5.438m-36.403 3.92c5.316-5.56 11.521-5.524 16.002-5.65m8.648 1.216c6.517 1.521 14.774.944 16.313-.856"/>`),
		g.Group(children),
	)
}