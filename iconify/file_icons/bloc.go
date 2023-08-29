package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bloc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M213.86 0L0 132.508v245.856L213.86 512L428 378.364V132.508L213.86 0zm156.022 155.36l-156.023 90.566L57.836 155.36L213.86 58.688l156.023 96.673zM49.888 173.82l153.993 89.387v183.731L49.888 350.711V173.82zm173.949 273.118v-183.73L377.83 173.82v176.89l-153.993 96.227z"/>`),
		g.Group(children),
	)
}