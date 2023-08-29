package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 6v20h9.586L16 29.414L19.414 26H29V6H3zm2 2h22v16h-8.414L16 26.586L13.414 24H5V8zm10 2v1.188c-1.156.417-2 1.521-2 2.814c0 1.645 1.355 3 3 3c.566 0 1 .434 1 1c0 .566-.434 1-1 1c-.566 0-1-.434-1-1h-2c0 1.293.844 2.394 2 2.812V22h2v-1.188c1.156-.418 2-1.521 2-2.814c0-1.645-1.355-3-3-3c-.566 0-1-.434-1-1c0-.566.434-1 1-1c.566 0 1 .434 1 1h2c0-1.293-.844-2.394-2-2.812V10h-2z"/>`),
		g.Group(children),
	)
}