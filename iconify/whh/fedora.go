package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fedora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 480q0 27-41 65t-107.5 73.5t-165 60.5T512 704t-198-25t-165-60.5T41 545T0 480q0-53 170-77q9-62 22-122.5T224 152t44.5-110T320 0q28 0 94.5 32T512 64q27 0 52-10t42-22t43-22t55-10q28 0 73 135t74 267q173 24 173 78z"/>`),
		g.Group(children),
	)
}