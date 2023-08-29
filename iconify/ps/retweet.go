package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M94 328q-7 0-12-5.5T77 310V170H2L104 58l103 112h-76v103h81l53 55H94zm368-114h-75V74q0-7-5-12.5T370 56H199l53 55h81v103h-76l103 113z"/>`),
		g.Group(children),
	)
}