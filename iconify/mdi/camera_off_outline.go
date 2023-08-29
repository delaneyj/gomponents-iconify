package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12c0 .54-.1 1.05-.26 1.54L15 11.78A3.057 3.057 0 0 0 12.22 9l-1.76-1.74c.49-.16 1-.26 1.54-.26c2.76 0 5 2.24 5 5M9.88 4h4.24l1.83 2H20v10.8l1.88 1.88c.08-.21.12-.44.12-.68V6c0-1.11-.89-2-2-2h-3.17L15 2H9L7.18 4L8.6 5.4L9.88 4m12.23 17.46l-1.27 1.27L18.11 20H4a2 2 0 0 1-2-2V6c0-.58.25-1.1.65-1.46L1.11 3l1.28-1.27l19.72 19.73M9 12c0 1.66 1.34 3 3 3c.33 0 .65-.07.94-.17l-3.77-3.77c-.1.3-.17.61-.17.94m7.11 6l-1.66-1.66c-.73.41-1.56.66-2.45.66c-2.76 0-5-2.24-5-5c0-.89.25-1.72.66-2.45L4.11 6H4v12h12.11Z"/>`),
		g.Group(children),
	)
}