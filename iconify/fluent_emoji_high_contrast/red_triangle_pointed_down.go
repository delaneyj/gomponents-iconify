package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.268 23c.77 1.333 2.694 1.333 3.464 0l6.062-10.5c.77-1.333-.192-3-1.732-3H9.938c-1.54 0-2.502 1.667-1.732 3L14.268 23ZM16 22L9.938 11.5h12.124L16 22Z"/>`),
		g.Group(children),
	)
}