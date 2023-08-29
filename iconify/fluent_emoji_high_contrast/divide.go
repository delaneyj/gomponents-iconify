package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.985 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM5.06 13.44C3.38 13.44 2 14.82 2 16.5c0 1.69 1.38 3.07 3.06 3.06h21.85c1.68 0 3.06-1.37 3.06-3.06c0-1.68-1.37-3.06-3.06-3.06H5.06Z"/>`),
		g.Group(children),
	)
}