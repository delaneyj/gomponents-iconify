package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steinberg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.48 12.63c.643-.372.643-1.01 0-1.38l-4.769-2.755c-.642-.37-1.195-.052-1.195.69v5.508c0 .742.553 1.06 1.195.69zM12 2.724a9.275 9.275 0 1 1-.001 18.55a9.275 9.275 0 0 1 0-18.55M12 0C5.383 0 .002 5.383.002 12s5.382 12 12 12S24 18.617 24 12S18.617 0 12 0z"/>`),
		g.Group(children),
	)
}