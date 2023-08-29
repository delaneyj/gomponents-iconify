package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownloadO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.1 9.8v-.6c0-2.4-1.9-4.3-4.2-4.3c-.3.1-.6.1-.9.1V2H7v2.4c-.4-.3-.9-.4-1.3-.4c-1.6 0-2.9 1.3-2.9 2.9c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.6C0 13.3 1.5 15 3.3 15h10.3c1.4 0 2.4-1.5 2.4-2.7s-.8-2.3-1.9-2.5zm-.5 4.2H3.3C2.1 14 1 12.7 1 11.4s1.1-2.6 2.3-2.6h.4l1.4.2l-.9-1c-.2-.3-.4-.7-.4-1.2c0-1 .8-1.8 1.8-1.8c.5 0 1 .2 1.3.6V8H5l3 4l3-4H9V6.1c.3-.1.6-.1.9-.1c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6l.8.1c.7 0 1.4.7 1.4 1.5c0 .7-.6 1.6-1.4 1.6z"/>`),
		g.Group(children),
	)
}