package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m670.312 0l177.246 177.295l-241.21 241.211l175.146 175.146l241.211-241.211L1200 529.688V0H670.312zM418.506 606.348L177.295 847.559L0 670.312V1200h529.688l-177.246-177.295l241.211-241.211l-175.147-175.146z"/>`),
		g.Group(children),
	)
}