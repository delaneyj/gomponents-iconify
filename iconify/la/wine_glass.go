package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m8.3 5l-.3.6C7.4 6.9 7 8.7 7 11c0 4.6 3.5 8.4 8 8.9V26h-5v2h12v-2h-5v-6.1c4.5-.5 8-4.3 8-8.9c0-2.3-.4-4.1-1-5.4l-.2-.6H8.3zm1.3 2h12.8c.4 1.2.6 2.6.6 4c0 2.1-.9 3.9-2.3 5.2l-.1.1c-.1.1-.2.2-.3.2c-.1.1-.2.1-.2.2c-.1.1-.2.1-.3.2c-.1.1-.2.1-.3.2c-.1 0-.2.1-.2.1c-.1.1-.2.1-.4.2c-.1 0-.2.1-.2.1c-.1.1-.3.1-.4.1c-.1 0-.1.1-.2.1s-.3.1-.4.1c-.1 0-.1 0-.2.1H16c-3.9 0-7-3.1-7-7V9.4L9.6 7z"/>`),
		g.Group(children),
	)
}