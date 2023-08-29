package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printbot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.42 15.125h15.16M16.42 18.72h15.16m-15.16-7.449h15.16M13.862 6.342v14.8H7.841a2.461 2.461 0 0 0-2.467 2.466v15.584a2.461 2.461 0 0 0 2.467 2.466h32.318a2.461 2.461 0 0 0 2.467-2.466V23.608a2.461 2.461 0 0 0-2.467-2.466h-6.021v-14.8Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.862 6.342h20.275v17.38H13.862z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.63 21.302v10.905a1.546 1.546 0 0 0 1.674 1.558c.827-.006 1.674 0 1.674 0H34.06s.77.026 1.673 0a1.717 1.717 0 0 0 1.674-1.573V21.179"/><circle cx="40.156" cy="35.763" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.382 41.658v-3.257c.058-.189.17-.465.367-.465H35.52a.58.58 0 0 1 .367.465v3.257"/>`),
		g.Group(children),
	)
}