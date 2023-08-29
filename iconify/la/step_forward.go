package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 5v10.5l-.438-.313l-13-9L7 5.095v21.812l1.563-1.093l13-9L22 16.5V27h2V5zM9 8.906L19.25 16L9 23.094z"/>`),
		g.Group(children),
	)
}