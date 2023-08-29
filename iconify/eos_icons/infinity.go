package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.385 8.609A5.469 5.469 0 0 0 18.5 7a5.568 5.568 0 0 0-3.928 1.609L12 11.092L9.428 8.61a5.5 5.5 0 1 0 0 7.782L12 13.83l2.572 2.562a5.514 5.514 0 1 0 7.813-7.782Zm-14.37 6.374a3.541 3.541 0 0 1-4.986 0a3.518 3.518 0 0 1 4.985-4.965l2.486 2.474Zm12.956 0a3.539 3.539 0 0 1-4.985 0l-2.486-2.49l2.486-2.476a3.541 3.541 0 0 1 4.985 0a3.506 3.506 0 0 1 0 4.965Z"/>`),
		g.Group(children),
	)
}