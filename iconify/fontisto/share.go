package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.237 15.638h-.003a4.04 4.04 0 0 0-3.004 1.334l-.003.004l-8.948-4.348c0-.167.084-.418.084-.669a1.425 1.425 0 0 0-.087-.595l.003.01l8.948-4.348a4.04 4.04 0 0 0 3.007 1.338h.004a4.181 4.181 0 1 0-4.181-4.181a1.425 1.425 0 0 0 .087.595l-.003-.01l-8.948 4.348a4.04 4.04 0 0 0-3.007-1.338h-.004a4.181 4.181 0 0 0 0 8.362h.003a4.04 4.04 0 0 0 3.004-1.334l.003-.004l8.948 4.348c0 .167-.084.418-.084.669a4.181 4.181 0 1 0 8.363-.092a4.09 4.09 0 0 0-4.09-4.09l-.097.001z"/>`),
		g.Group(children),
	)
}