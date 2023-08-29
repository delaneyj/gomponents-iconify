package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shopware(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.83 3 3 8.83 3 16s5.83 13 13 13c2.9 0 5.58-.96 7.75-2.57c-.03-.02-.07-.04-.12-.07c-1.12-.84-2.76-1.48-4.5-2.14C14.76 22.52 10 20.61 10 15c0-5.24 4.8-8 11.49-8c.97 0 1.96.06 2.93.19l.29.04c.37.08.74.16 1.1.25C23.43 4.74 19.91 3 16 3zm5.25 9C16.81 12 17 14.94 17 15.01c0 2.4 2.75 3.49 5.95 4.76c1.55.62 3.16 1.25 4.65 2.1c.89-1.76 1.4-3.76 1.4-5.87c0-.18 0-.36-.02-.54c-.03-.02-.07-.04-.1-.07c-2.74-2.41-4.94-3.39-7.63-3.39z"/>`),
		g.Group(children),
	)
}