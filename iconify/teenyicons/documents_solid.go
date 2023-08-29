package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1.5A1.5 1.5 0 0 1 4.5 0h6.151L14 2.232V11.5a1.5 1.5 0 0 1-1.5 1.5H12v.5a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 1 13.5v-10A1.5 1.5 0 0 1 2.5 2H3v-.5ZM3 3h-.5a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5V13H4.5A1.5 1.5 0 0 1 3 11.5V3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}