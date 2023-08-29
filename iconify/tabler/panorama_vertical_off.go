package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaVerticalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 3h10.53c.693 0 1.18.691.935 1.338c-1.098 2.898-1.573 5.795-1.425 8.692m.828 4.847c.172.592.37 1.185.595 1.778A1 1 0 0 1 17.529 21h-11c-.692 0-1.208-.692-.962-1.34c1.697-4.486 1.903-8.973.619-13.46M3 3l18 18"/>`),
		g.Group(children),
	)
}