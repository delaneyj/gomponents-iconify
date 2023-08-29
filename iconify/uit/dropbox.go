package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.3 12.5l-3.6-3.3l3.6-2.9c.1 0 .1-.1.1-.1c.1-.2 0-.5-.2-.7l-5-2.7c-.2-.1-.4-.1-.5 0L12 5.6L8.3 2.8c-.2-.1-.4-.1-.5 0l-5 2.7c-.1 0-.1.1-.1.1c-.3.1-.2.5 0 .6l3.6 2.9l-3.6 3.3l-.1.1c-.1.2 0 .5.2.7l3.7 2v2.5c0 .2.1.3.2.4l5 3c.1 0 .2.1.3.1c.1 0 .2 0 .3-.1l5-3c.2-.1.2-.3.2-.4v-2.5l3.7-2l.1-.1c.3-.1.2-.4 0-.6zM16 3.7l4.1 2.2L17 8.6l-4.1-2.4L16 3.7zm.1 5.5L12 11.9L7.9 9.2L12 6.8l4.1 2.4zM3.9 5.9L8 3.7l3.2 2.4L7 8.6L3.9 5.9zm0 6.8l3.2-3l4.1 2.7L8 15l-4.1-2.3zm12.6 4.9L12 20.3l-4.5-2.7v-1.7l.3.1H8c.1 0 .2 0 .3-.1L12 13l3.7 2.9c.1.1.2.1.3.1c.1 0 .2 0 .2-.1l.3-.1v1.8zM16 15l-3.2-2.5l4.1-2.7l3.2 3L16 15z"/>`),
		g.Group(children),
	)
}