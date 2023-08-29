package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoystickSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 3.5a3.5 3.5 0 1 1 4 3.465V10h5.5a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-12a.5.5 0 0 1-.5-.5v-4a.5.5 0 0 1 .5-.5H7V6.965A3.5 3.5 0 0 1 4 3.5Z"/><path fill="currentColor" d="M3 8v1h2V8H3Z"/>`),
		g.Group(children),
	)
}