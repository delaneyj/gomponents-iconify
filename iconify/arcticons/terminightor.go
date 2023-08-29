package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminightor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M24 9.2A16.85 16.85 0 1 1 7.1 26.1A16.94 16.94 0 0 1 24 9.2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.31 19.3l-5.5 1.5m-17.1-1.5l5.5 1.5M17.68 5c-5.45-3.29-12 .8-11.1 7.5c.1.8 1.57 1.78 2.4.8a17.48 17.48 0 0 1 9.1-5.7c1.49-.36.41-2.11-.4-2.6Zm12.64 0c5.45-3.29 12 .8 11.1 7.5c-.1.8-1.57 1.78-2.4.8a17.48 17.48 0 0 0-9.1-5.7c-1.49-.36-.41-2.11.4-2.6ZM24 13.4v12.65l4.5 6"/>`),
		g.Group(children),
	)
}