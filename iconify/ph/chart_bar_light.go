package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 202h-10V40a6 6 0 0 0-6-6h-56a6 6 0 0 0-6 6v42H96a6 6 0 0 0-6 6v42H48a6 6 0 0 0-6 6v66H32a6 6 0 0 0 0 12h192a6 6 0 0 0 0-12ZM158 46h44v156h-44Zm-56 48h44v108h-44Zm-48 48h36v60H54Z"/>`),
		g.Group(children),
	)
}