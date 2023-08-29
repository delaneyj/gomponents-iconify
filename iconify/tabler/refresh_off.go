package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 11A8.1 8.1 0 0 0 8.729 4.695m-2.41 1.624A8.083 8.083 0 0 0 4.5 9M4 5v4h4m-4 4a8.1 8.1 0 0 0 13.671 4.691M20 16v-1h-1M3 3l18 18"/>`),
		g.Group(children),
	)
}