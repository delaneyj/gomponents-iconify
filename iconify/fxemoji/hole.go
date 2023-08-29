package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<ellipse cx="256" cy="249.721" fill="#464A4C" rx="219.043" ry="75.527"/><path fill="#132028" d="M75.169 287.933c0-41.713 98.069-75.527 219.042-75.527c74.839 0 140.899 12.944 180.41 32.691c-6.932-39.558-102.154-70.903-218.622-70.903c-120.974 0-219.042 33.815-219.042 75.527c0 15.908 14.275 30.664 38.632 42.837a26.867 26.867 0 0 1-.42-4.625z"/>`),
		g.Group(children),
	)
}