package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardPulse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1h-4zM6 1h4v2H6V1zm7 14H3V3h2v1h6V3h2v12z"/><path fill="currentColor" d="M9.3 13c-.2 0-.3-.1-.4-.3l-.8-4.8l-.7 3.1c0 .1-.1.2-.3.3c-.1 0-.3 0-.4-.1l-1-1.3H4.4c-.2 0-.4-.2-.4-.4s.2-.4.4-.4H6c.1 0 .2.1.3.1l.6.8l.9-4.3c0-.2.2-.3.4-.3s.3.2.3.4l.9 5.3l.6-1.7c.1-.1.2-.2.3-.2h1.3c.2 0 .4.2.4.4s-.2.4-.4.4h-1l-1 2.9s-.2.1-.3.1z"/>`),
		g.Group(children),
	)
}