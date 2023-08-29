package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraLensFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.827 21.763L14.31 14l3.532 6.117A9.955 9.955 0 0 1 12 22c-.746 0-1.473-.082-2.173-.237ZM7.89 21.12A10.029 10.029 0 0 1 2.458 15h8.965L7.89 21.119ZM2.05 13a9.964 9.964 0 0 1 2.583-7.761L9.112 13H2.05Zm4.109-9.117A9.955 9.955 0 0 1 12 2c.746 0 1.473.082 2.173.237L9.69 10L6.159 3.883ZM16.11 2.88A10.028 10.028 0 0 1 21.542 9h-8.965l3.533-6.119ZM21.95 11a9.964 9.964 0 0 1-2.583 7.761L14.888 11h7.064Z"/>`),
		g.Group(children),
	)
}