package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.406 5.2a3.5 3.5 0 0 0-2.753 1.338l-3.561 4.535a1.5 1.5 0 0 0 0 1.853l3.561 4.536A3.5 3.5 0 0 0 8.406 18.8h10.647a2.5 2.5 0 0 0 2.5-2.5V7.7a2.5 2.5 0 0 0-2.5-2.5H8.406Zm2.064 3.27a.75.75 0 0 1 1.06 0L14 10.94l2.47-2.47a.75.75 0 1 1 1.06 1.06L15.06 12l2.47 2.47a.75.75 0 0 1-1.06 1.06L14 13.06l-2.47 2.47a.75.75 0 1 1-1.06-1.06L12.94 12l-2.47-2.47a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}