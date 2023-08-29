package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aws(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m2 87l204 67v292L3 377zm255 360l205-70V89l-204 64zm-25-328l188-61L232 1L39 58z"/>`),
		g.Group(children),
	)
}