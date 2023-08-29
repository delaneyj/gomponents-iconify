package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.25 19.25h-.66c-.38-1.36-1.55-4.74-4.4-7.25c2.85-2.51 4.02-5.89 4.4-7.25h.66c.41 0 .75-.34.75-.75s-.34-.75-.75-.75H5.75c-.41 0-.75.34-.75.75s.34.75.75.75h.66c.38 1.36 1.55 4.74 4.4 7.25c-2.85 2.51-4.02 5.89-4.4 7.25h-.66c-.41 0-.75.34-.75.75s.34.75.75.75h12.5c.41 0 .75-.34.75-.75s-.34-.75-.75-.75ZM7.98 4.75h8.03c-.45 1.42-1.6 4.25-4.02 6.28c-2.42-2.04-3.57-4.86-4.02-6.28ZM12 12.97c2.42 2.04 3.57 4.86 4.02 6.28H7.98c.45-1.42 1.6-4.25 4.02-6.28Z"/>`),
		g.Group(children),
	)
}