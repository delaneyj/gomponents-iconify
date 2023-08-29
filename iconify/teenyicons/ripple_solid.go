package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RippleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m4.193 5.155l-3.06-3.316l.734-.678l3.061 3.316a3.5 3.5 0 0 0 5.144 0l3.06-3.316l.735.678l-3.06 3.316a4.5 4.5 0 0 1-6.614 0Zm5.879 5.368a3.5 3.5 0 0 0-5.144 0l-3.06 3.316l-.735-.678l3.06-3.316a4.5 4.5 0 0 1 6.614 0l3.06 3.316l-.734.678l-3.061-3.316Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}