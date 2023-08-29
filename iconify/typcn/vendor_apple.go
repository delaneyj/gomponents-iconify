package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VendorApple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.9 6.6s-.1-1.6.9-3L15.6 2s.1 1.6-.9 3l-2.8 1.6zm5.4 5.6c0-1.5.8-2.8 2-3.6l-.9-.9c-.5-.3-1.1-.7-2.4-.7c-1.4 0-2.4.9-3.7.9c-1.3 0-2.2-.8-3.1-.9c-.7 0-1.4 0-2.1.3c-.5.2-1.2.7-1.6 1.2c-.6.6-1.2 1.9-1.3 3.1c-.1 1.2-.1 2.1.2 3.2c.2.9.6 1.8 1 2.6c.3.6.6 1.2 1 1.8c.3.4.7.8 1.1 1.1c.3.2.6.4 1 .6c.2 0 .5.1.8.1c.6-.1 1.6-.9 2.4-1.1c.4-.1.8-.1 1.3 0c.7.1 1.4.9 2.2 1c.6.1 1.2 0 1.7-.3c.4-.2.7-.5 1-.9c.4-.4.7-.9 1-1.3c.4-.7.9-1.5 1.1-2.3c-1.6-.6-2.7-2.1-2.7-3.9z"/>`),
		g.Group(children),
	)
}