package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 9v2s-.8 1.7-4 1.9V6h2.2c.2.3.5.5.8.5c.6 0 1-.4 1-1s-.4-1-1-1c-.4 0-.7.2-.8.5H9V3.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2S6 .9 6 2c0 .7.4 1.4 1 1.7V5H4.8c-.1-.3-.4-.5-.8-.5c-.6 0-1 .4-1 1s.4 1 1 1c.4 0 .7-.2.8-.5H7v7c-3.3-.3-4-2-4-2V9H0s2.8 7 8 7c5 0 8-7 8-7h-3zM8 1c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1z"/>`),
		g.Group(children),
	)
}