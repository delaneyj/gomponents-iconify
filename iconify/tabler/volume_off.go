package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 8a5 5 0 0 1 1.912 4.934m-1.377 2.602A5 5 0 0 1 15 16m2.7-11a9 9 0 0 1 2.362 11.086m-1.676 2.299A9 9 0 0 1 17.7 19M9.069 5.054L9.5 4.5A.8.8 0 0 1 11 5v2m0 4v8a.8.8 0 0 1-1.5.5L6 15H4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h2l1.294-1.664M3 3l18 18"/>`),
		g.Group(children),
	)
}