package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WateringSoil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 12h2m16 0h2M3 20.01l.01-.011M6 16.01l.01-.011M9 20.01l.01-.011m5.99.011l.01-.011M18 16.01l.01-.011M21 20.01l.01-.011M12.396 3.396L15.5 6.5a4.95 4.95 0 1 1-7 0l3.104-3.104a.56.56 0 0 1 .792 0Z"/>`),
		g.Group(children),
	)
}