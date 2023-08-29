package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileVueFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m213.66 82.34l-56-56A8 8 0 0 0 152 24H56a16 16 0 0 0-16 16v72a8 8 0 0 0 8 8h160a8 8 0 0 0 8-8V88a8 8 0 0 0-2.34-5.66ZM152 88V44l44 44Zm-64.46 66.69l-20 56a8 8 0 0 1-15.07 0l-20-56a8 8 0 0 1 15.07-5.38L60 184.21l12.47-34.9a8 8 0 0 1 15.07 5.38ZM184 160v12h16a8 8 0 0 1 0 16h-16v12h24a8 8 0 0 1 0 16h-32a8 8 0 0 1-8-8v-56a8 8 0 0 1 8-8h32a8 8 0 0 1 0 16Zm-32-8v38a26 26 0 0 1-52 0v-38a8 8 0 0 1 16 0v38a10 10 0 0 0 20 0v-38a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}