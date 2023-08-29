package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.001 1.993C6.486 1.994 2 6.48 2 11.994c.001 5.514 4.487 10 10 10c5.515 0 10.001-4.486 10.001-10s-4.486-10-10-10.001zM12 19.994c-4.41 0-7.999-3.589-8-8c0-4.411 3.589-8 8.001-8.001c4.411.001 8 3.59 8 8.001s-3.589 8-8.001 8z"/><path fill="currentColor" d="m12.001 8.001l-4.005 4.005h3.005V16h2v-3.994h3.004z"/>`),
		g.Group(children),
	)
}