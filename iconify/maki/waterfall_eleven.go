package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.9 2H11V1H4C2.4 1 1 2.2 1 3.9v3.4C0 7.8-.3 9 .3 10c.6 1 1.8 1.3 2.7.7c.7.4 1.7.3 2.3-.2c1.4.9 3.2.6 4.2-.8c.9-1.4.6-3.2-.8-4.2c-.2-.1-.5-.2-.7-.3V4c0-1.1.8-2 1.9-2zM9 8c0 1.1-.9 2-2 2c-1 0-1.1-.3-1.3-.5h-.8c-.2.2-.4.5-.9.5s-.7-.2-.8-.5h-.4c-.1.3-.4.5-.8.5c-.6 0-1-.4-1-1s.4-1 1-1V5s0-.5.5-.5s.5.5.5.5v2.5s0 .5.5.5s.5-.5.5-.5V6s0-.5.5-.5s.5.5.5.5v1.5s0 .5.5.5s.5-.5.5-.5V5s0-.5.5-.5s.5.5.5.5v1c1.1 0 2 .9 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}