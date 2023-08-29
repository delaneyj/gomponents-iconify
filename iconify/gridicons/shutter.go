package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.9 4.8s-.7 5.6-3.5 10.2c1.7-.3 3.9-.9 6.6-2c0 0 .7-4.6-3.1-8.2zm-6 2.8c-1.1-1.3-2.7-3-5-4.7C5.1 4.2 3 6.6 2.3 9.6C7 7.7 11 7.5 12.9 7.6zm3.4 2.9c.6-1.6 1.2-3.9 1.6-6.7c-4.1-3-8.6-1.5-8.6-1.5s4.4 3.4 7 8.2zm-5.2 6c1.1 1.3 2.7 3 5 4.7c0 0 4.3-1.6 5.6-6.7c0-.1-5.3 2.1-10.6 2zm-3.4-3.1c-.6 1.6-1.2 3.8-1.5 6.7c0 0 3.6 2.9 8.6 1.5c0 0-4.6-3.4-7.1-8.2zM2 11.1s-.7 4.5 3.1 8.2c0 0 .7-5.7 3.5-10.3c-1.7.3-4 .9-6.6 2.1z"/>`),
		g.Group(children),
	)
}