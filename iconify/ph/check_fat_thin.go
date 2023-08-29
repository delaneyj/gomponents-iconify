package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckFatThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m240.47 71.09l-24-23.58a12 12 0 0 0-17 0l-92.71 91.34a4 4 0 0 1-5.66 0v-.06l-36.62-35.28a12 12 0 0 0-17 0l-24 24a12 12 0 0 0 0 17l71.61 72a12 12 0 0 0 17 0l128.4-128.43a12 12 0 0 0-.02-16.99Zm-5.63 11.34l-128.41 128.4a4 4 0 0 1-5.65 0l-71.61-72a4 4 0 0 1 0-5.66l24-24A4 4 0 0 1 56 108a4.09 4.09 0 0 1 2.9 1.21l36.66 35.29a12 12 0 0 0 16.93 0l92.71-91.33a4 4 0 0 1 5.68 0l24 23.58a4 4 0 0 1-.04 5.68Z"/>`),
		g.Group(children),
	)
}