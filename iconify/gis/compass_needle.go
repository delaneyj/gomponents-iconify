package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassNeedle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50.03 5a2.516 2.516 0 0 0-2.43 1.76L34.493 48.548a2.51 2.51 0 0 0-.372 1.454c-.026.51.104 1.017.372 1.452l13.105 41.782c.737 2.352 4.065 2.352 4.802 0l13.105-41.785c.27-.436.399-.945.372-1.456a2.513 2.513 0 0 0-.372-1.45L52.401 6.76A2.513 2.513 0 0 0 50.03 5ZM39.403 50.288h6.205c.152 2.306 2.048 4.134 4.392 4.134c2.344 0 4.24-1.828 4.392-4.134h6.461L50 84.078Z" color="currentColor"/>`),
		g.Group(children),
	)
}