package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M175.736 1024.264c234.315 234.314 614.213 234.314 848.527 0s234.314-614.212 0-848.527c-234.314-234.314-614.212-234.314-848.527 0c-234.314 234.314-234.314 614.212 0 848.527zm94.931-65.618l-29.261-29.262l261.125-426.853l426.853-261.125l29.261 29.261l-261.124 426.854l-426.854 261.125z"/>`),
		g.Group(children),
	)
}