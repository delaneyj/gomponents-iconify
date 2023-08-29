package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 200a12 12 0 0 1-24 0c0-77.2-62.8-140-140-140a12 12 0 0 1 0-24c90.43 0 164 73.57 164 164ZM56 108a12 12 0 0 0 0 24a68.07 68.07 0 0 1 68 68a12 12 0 0 0 24 0a92.1 92.1 0 0 0-92-92Zm4 72a16 16 0 1 0 16 16a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}