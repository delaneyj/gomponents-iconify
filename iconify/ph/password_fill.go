package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PasswordFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 48H40a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM64 168a8 8 0 0 1-16 0V88a8 8 0 0 1 16 0Zm73.3-39.81l-12.36 4l7.64 10.5a8 8 0 1 1-13 9.41L112 141.61l-7.63 10.51a8 8 0 1 1-13-9.41l7.64-10.5l-12.36-4a8 8 0 1 1 5-15.21l12.35 4v-13a8 8 0 0 1 16 0v13l12.35-4a8 8 0 1 1 5 15.21Zm72 0l-12.36 4l7.64 10.5a8 8 0 1 1-13 9.41L184 141.61l-7.63 10.51a8 8 0 1 1-13-9.41l7.64-10.5l-12.36-4a8 8 0 1 1 5-15.21l12.35 4v-13a8 8 0 0 1 16 0v13l12.35-4a8 8 0 0 1 5 15.21Z"/>`),
		g.Group(children),
	)
}