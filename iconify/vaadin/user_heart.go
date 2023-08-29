package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.2 16h.6v-.6l-.6.6zm-5.6-2.1c-.7-.7-1-1.8-.8-2.8S8.6 9.3 9.5 9c0-.1-.1-.2-.1-.2c0-.6.8-.4 1.4-1.5c0 0 2.7-7.3-2.9-7.3c-5.5 0-2.8 7.3-2.8 7.3c.6 1 1.4.8 1.4 1.5s-.7.7-1.4.8C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h9.6l-2.1-2.1z"/><path fill="currentColor" d="M14.9 10.1c-.2-.1-.5-.1-.7-.1c-.7 0-1.3.6-1.7 1.1c-.4-.5-1-1.1-1.7-1.1c-.3 0-.5 0-.7.1c-1.2.4-1.4 2-.5 2.9l3 2.9l3-2.9c.8-.9.5-2.5-.7-2.9z"/>`),
		g.Group(children),
	)
}