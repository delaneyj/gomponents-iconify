package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 22h-4v-1h4v1m-1-12V7h-2v3l-1 1.5V20h4v-8.5L13 10m-1-8s-1 1-1 3v1h2V5s0-2-1-3Z"/>`),
		g.Group(children),
	)
}