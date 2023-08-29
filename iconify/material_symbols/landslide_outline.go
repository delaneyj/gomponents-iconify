package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandslideOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22h20l-6-8l-5-2l-3-4H2v14Zm2-2v-1.6l2 .65l9.025-3L18 20H4Zm2-3.05l-2-.675V14.4l2 .65l3.95-1.3l2.4 1.075L6 16.95ZM18.5 14l4.5-2V8l-4.5-1L16 9v3l2.5 2ZM6 12.95l-2-.675V10h3l1.625 2.075L6 12.95Zm12.8-1.275l-.8-.625v-1.1l1-.8l2 .45v1.1l-2.2.975ZM12 8l5-2V1l-5-1l-3 2v4l3 2Zm.225-2.25L11 4.925v-1.85l1.425-.95L15 2.65v2l-2.775 1.1Z"/>`),
		g.Group(children),
	)
}