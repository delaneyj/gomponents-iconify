package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#FF473E" d="m181.816 496.127l68.178-102.898c2.849-4.3 9.162-4.3 12.011 0l68.178 102.898c6.382 9.633 21.383 5.114 21.383-6.441V16.662c0-6.441-5.221-11.662-11.662-11.662h-167.81c-6.441 0-11.662 5.221-11.662 11.662v473.024c.001 11.555 15.002 16.074 21.384 6.441z"/>`),
		g.Group(children),
	)
}