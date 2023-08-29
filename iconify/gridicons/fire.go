package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 13.4c0 1.2-.2 2.3-.7 3.4c-.5 1.1-1.2 2-2.1 2.9c-.9.8-1.9 1.5-3.1 1.9c-.5.2-1.1.4-1.7.5l.3-.1s2.2-1.3 2.7-4.1c.2-2.5-1.9-3.6-1.9-3.6s-1.1 1.4-2.6 1.4c-2.4 0-2-4-2-4s-3.6 1.9-3.6 5.9c0 2.5 2.6 4.4 2.6 4.4c-.3-.1-.7-.2-1-.3c-1.2-.4-2.2-1.1-3.1-1.9c-.9-.8-1.6-1.8-2.1-2.9c-.5-1.1-.7-2.2-.7-3.4c0-3.9 4.6-7.7 4.6-7.7s.4 3 2.7 3c4.2 0 2.7-6.7 2.7-6.7s3.8 2.2 4.9 8.2c2.1-.3 1.9-3 1.9-3s2.2 2.9 2.2 6.1"/>`),
		g.Group(children),
	)
}