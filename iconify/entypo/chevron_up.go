package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.484 12.452c-.436.446-1.043.481-1.576 0L10 8.705l-3.908 3.747c-.533.481-1.141.446-1.574 0c-.436-.445-.408-1.197 0-1.615c.406-.418 4.695-4.502 4.695-4.502a1.095 1.095 0 0 1 1.576 0s4.287 4.084 4.695 4.502c.409.418.436 1.17 0 1.615z"/>`),
		g.Group(children),
	)
}