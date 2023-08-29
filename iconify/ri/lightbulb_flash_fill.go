package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbFlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.941 18c-.297-1.273-1.637-2.314-2.187-3a8 8 0 1 1 12.49.002c-.55.685-1.888 1.726-2.185 2.998H7.941ZM16 20v1a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-1h8Zm-3-9.995V6l-4.5 6.005H11v4l4.5-6H13Z"/>`),
		g.Group(children),
	)
}