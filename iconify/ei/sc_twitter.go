package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScTwitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M39.2 16.8c-1.1.5-2.2.8-3.5 1c1.2-.8 2.2-1.9 2.7-3.3c-1.2.7-2.5 1.2-3.8 1.5c-1.1-1.2-2.7-1.9-4.4-1.9c-3.3 0-6.1 2.7-6.1 6.1c0 .5.1.9.2 1.4c-5-.2-9.5-2.7-12.5-6.3c-.5.7-.8 1.7-.8 2.8c0 2.1 1.1 4 2.7 5c-1 0-1.9-.3-2.7-.8v.1c0 2.9 2.1 5.4 4.9 5.9c-.5.1-1 .2-1.6.2c-.4 0-.8 0-1.1-.1c.8 2.4 3 4.2 5.7 4.2c-2.1 1.6-4.7 2.6-7.5 2.6c-.5 0-1 0-1.4-.1c2.4 1.9 5.6 2.9 9 2.9c11.1 0 17.2-9.2 17.2-17.2V20c1.2-.9 2.2-1.9 3-3.2z"/>`),
		g.Group(children),
	)
}