package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clickup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.4 14c-.1-.1-.2-.2-.4-.3c-.1.1-.2.2-.3.2C20.1 17 16.6 20 13 23.1q-.5.4-.9-.1c-1.5-1.8-3-3.5-4.5-5.2c-.3-.3-.3-.5 0-.8l16-13.6c.4-.3.6-.4 1 0l15.8 13.8c.3.3.3.4.1.7l-4.2 4.8c-.2.2-.4.5-.6.7c-.2.3-.3.2-.5 0L24.4 14Zm-.3 30.9c-5.2 0-9.7-2-13.6-5.4C9.2 38.3 8 37 6.8 35.6c-.3-.4-.2-.5.1-.8c1.9-1.4 3.8-2.9 5.7-4.3c.3-.2.4-.2.6.1c1.3 1.8 2.9 3.3 4.8 4.5c1.3.8 2.8 1.4 4.4 1.6c3 .4 5.7-.3 8.2-2c1.6-1.1 2.9-2.5 4-4c.2-.3.4-.3.7-.1l5.7 4.2c.3.2.4.3.2.6c-2.8 3.7-6.2 6.7-10.6 8.3c-2 .8-4.1 1.2-6.5 1.2Z"/>`),
		g.Group(children),
	)
}