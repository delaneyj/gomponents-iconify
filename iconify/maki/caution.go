package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.093 11.892L6.84 1.39a.752.752 0 0 1 1.32 0l5.747 10.502a.75.75 0 0 1-.66 1.11H1.753a.75.75 0 0 1-.66-1.11ZM8.3 8l.403-2.418A.5.5 0 0 0 8.21 5H6.79a.5.5 0 0 0-.493.582L6.7 8h1.6Zm.3 1.9a1.1 1.1 0 1 0-2.2 0a1.1 1.1 0 0 0 2.2 0Z"/>`),
		g.Group(children),
	)
}