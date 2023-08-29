package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prohibited(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.95 2C8.26 2 2 8.26 2 15.95C2 23.64 8.26 29.9 15.95 29.9c7.69 0 13.95-6.25 13.95-13.95C29.9 8.25 23.65 2 15.95 2Zm10.94 13.95c0 2.47-.83 4.75-2.22 6.58L9.37 7.24c1.83-1.39 4.11-2.22 6.58-2.22c6.03-.01 10.94 4.9 10.94 10.93Zm-21.88 0c0-2.48.84-4.76 2.23-6.59l15.3 15.3a10.906 10.906 0 0 1-6.59 2.23c-6.03 0-10.94-4.91-10.94-10.94Z"/>`),
		g.Group(children),
	)
}