package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserVisible(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM8 16a4 4 0 0 0-4 4h7.05v2H2v-2a6 6 0 0 1 6-6h3v2H8Zm8.248 3v-2h2.5v2h-2.5Z"/><path fill="currentColor" d="M17.5 22.5c4.418 0 6.09-4.5 6.09-4.5s-1.674-4.5-6.09-4.5c-4.417 0-6.09 4.5-6.09 4.5s1.671 4.5 6.09 4.5Zm-.002-2c-2.616 0-3.87-2.5-3.87-2.5s1.25-2.5 3.87-2.5s3.87 2.5 3.87 2.5s-1.254 2.5-3.87 2.5Z"/>`),
		g.Group(children),
	)
}