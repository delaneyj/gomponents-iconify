package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/><path fill="currentColor" d="M8 6C5.5 6 6.7 9.2 6.7 9.2c.3.4.7.4.7.6c0 .3-.3.3-.6.4c-.5.1-.9-.1-1.4.8c-.3.4-.4 2-.4 2h6s-.1-1.6-.4-2c-.4-.8-.9-.7-1.4-.8c-.3 0-.6-.1-.6-.4s.3-.2.6-.6C9.3 9.2 10.5 6 8 6z"/>`),
		g.Group(children),
	)
}