package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluebot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M30.452 24a5.25 5.25 0 0 1 0 10.5h-5.347v-21h5.347a5.25 5.25 0 0 1 0 10.5Zm0 0h-5.347m-7.557 0a5.25 5.25 0 1 0 0 10.5h5.347v-21h-5.347a5.25 5.25 0 0 0 0 10.5Zm0 0h5.347"/>`),
		g.Group(children),
	)
}