package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FStopSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.055 4.105a1.917 1.917 0 0 1 2.534-1.494l.41.146a.75.75 0 1 0 .503-1.413l-.41-.146A3.417 3.417 0 0 0 7.575 3.86l-.476 2.89H4.5a.75.75 0 1 0 0 1.5h2.352l-.434 2.632a1.917 1.917 0 0 1-2.924 1.303l-.34-.218a.75.75 0 1 0-.808 1.264l.34.217c2.067 1.323 4.814.1 5.213-2.322l.473-2.877H10.5a.75.75 0 0 0 0-1.5H8.62l.435-2.645Z"/>`),
		g.Group(children),
	)
}