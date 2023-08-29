package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlRice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19.66V21a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1.34A10 10 0 0 0 22 11a1 1 0 0 0-1-1a3.58 3.58 0 0 0-1.8-3a3.66 3.66 0 0 0-3.63-3.13a3.86 3.86 0 0 0-1 .13a3.7 3.7 0 0 0-5.11 0a3.86 3.86 0 0 0-1-.13A3.66 3.66 0 0 0 4.81 7A3.58 3.58 0 0 0 3 10a1 1 0 0 0-1 1a10 10 0 0 0 5 8.66zm-.89-11l.83-.26l-.16-.9a1.64 1.64 0 0 1 1.66-1.62a1.78 1.78 0 0 1 .83.2l.81.45l.5-.77a1.71 1.71 0 0 1 2.84 0l.5.77l.81-.45a1.78 1.78 0 0 1 .83-.2a1.65 1.65 0 0 1 1.67 1.6l-.16.85l.82.28A1.59 1.59 0 0 1 19 10H5a1.59 1.59 0 0 1 1.11-1.39zM19.94 12a8 8 0 0 1-4.39 6.16a1 1 0 0 0-.55.9V20H9v-.94a1 1 0 0 0-.55-.9A8 8 0 0 1 4.06 12z"/>`),
		g.Group(children),
	)
}