package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 4h20M3 8.01l.01-.011M3 16.01l.01-.011M6 12.01l.01-.011M6 20.01l.01-.011M9 8.01l.01-.011M9 16.01l.01-.011M12 12.01l.01-.011M12 20.01l.01-.011M15 8.01l.01-.011M15 16.01l.01-.011M18 12.01l.01-.011M18 20.01l.01-.011M21 8.01l.01-.011M21 16.01l.01-.011"/>`),
		g.Group(children),
	)
}