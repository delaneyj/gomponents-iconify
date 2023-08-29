package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.6 2.4c-.8-.8-2.1-.1-2.7.5l-3.5 3.8l-2.4-1l.6-.6c.2-.2.2-.5 0-.7c-.1-.2-.4-.1-.6 0l-.9.9l-1.8-.8l.5-.5c.2-.2.2-.5 0-.6c-.2-.2-.5-.2-.7 0l-.8.8l-.5-.2c-.8-.4-1.7-.3-2.3.4l5.8 5.8L6 12.6l-3.5-.2l-.5.7l3.1 1.8L6.9 18l.7-.5l-.2-3.5l2.5-2.3l5.8 5.8c.6-.6.8-1.6.4-2.3l-.2-.5l.8-.8c.2-.2.2-.5 0-.7c-.2-.2-.5-.2-.7 0l-.5.5l-.8-1.9l.9-.9c.2-.2.2-.5 0-.7c-.2-.2-.5-.2-.6 0l-.6.6l-1-2.4L17.2 5c.6-.5 1.2-1.9.4-2.6z"/>`),
		g.Group(children),
	)
}