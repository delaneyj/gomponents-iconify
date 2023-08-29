package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flowup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1012 574L818 392q-20-23-49 12v108H640q-26 0-45-19t-19-45V256q0-70-19.5-121T501 55t-80.5-42T320 0H0v192h320q27 0 45.5 18.5T384 256v192q0 70 20 120.5t55.5 79.5t80.5 42.5T640 704h129v107q29 35 49 12l194-182q12-14 12-33.5t-12-33.5z"/>`),
		g.Group(children),
	)
}