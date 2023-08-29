package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zulip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.54 10.22a6.27 6.27 0 0 1-2.2 4.78L26.21 26.76c-.21.2-.61-.1-.41-.4l4.82-9.63c.1-.3 0-.6-.3-.6H11.67c-2.91 0-5.21-2.6-5.21-5.81s2.3-5.82 5.21-5.82h24.56c2.91-.1 5.31 2.51 5.31 5.72ZM11.67 43.5h24.66c2.91 0 5.21-2.61 5.21-5.81s-2.3-5.82-5.21-5.82H17.68a.37.37 0 0 1-.3-.6l4.82-9.62a.29.29 0 0 0-.41-.41L8.66 32.87a6.3 6.3 0 0 0-2.2 4.82c0 3.2 2.4 5.81 5.21 5.81Z"/>`),
		g.Group(children),
	)
}