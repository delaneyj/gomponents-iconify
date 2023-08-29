package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceAnimalThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.052 17s-2.708 0-4.393 2.973C2.125 20.915 2 22.02 2 23.101V24v-4.995a7.464 7.464 0 0 1 6.444-7.394l.626-.086A8.606 8.606 0 0 0 16.5 3A2.5 2.5 0 0 0 19 5.5h4.5v1a3 3 0 0 1-3 3h-1v10.792c0 .875 0 1.458.45 2.333c0 0 .45.875 1.05.875M7.052 17l3.657.671A20 20 0 0 0 14.32 18h1.18m-8.448-1s-1.266 1.224-1.88 3.648c-.137.54-.172 1.099-.172 1.656V24m10.5-6v2.292c0 .875 0 1.458.45 2.333c0 0 .45.875 1.05.875M15.5 18l-.234-.312C13.613 15.483 12.443 12.702 12.9 10m2.599 8c2.564-.56 4-2.5 4-4a6.243 6.243 0 0 1-5.923-4.27L13.5 9.5c-2.5-1-5-4-5-9.5m-8 8v3.5a3 3 0 0 0 3 3h.013"/>`),
		g.Group(children),
	)
}