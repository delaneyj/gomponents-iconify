package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUploadO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.1 10.9v-.6c0-2.4-1.9-4.3-4.2-4.3c-.3 0-.6 0-.9.1V4h2L8 0L5 4h2v1.5c-.4-.2-.9-.3-1.3-.3c-1.6 0-2.9 1.2-2.9 2.8c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.5C0 14.3 1.5 16 3.3 16h10.3c1.4 0 2.4-1.4 2.4-2.6s-.8-2.2-1.9-2.5zm-.5 4.1H3.3C2.1 15 1 13.8 1 12.5S2.1 10 3.3 10h.4l1.3.3l-.8-1.2c-.2-.3-.4-.7-.4-1.1c0-1 .8-1.8 1.8-1.8c.5 0 1 .2 1.3.6V10h2V7.2c.3-.1.6-.1.9-.1c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6h.8c.7 0 1.4.7 1.4 1.5c.1.7-.5 1.6-1.3 1.6z"/>`),
		g.Group(children),
	)
}