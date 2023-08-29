package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attention(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.04 8l-3.174 5.5a1 1 0 0 0 .866 1.5h2.173a1 1 0 0 0 .724-.31L13 8L6.629 1.31A1 1 0 0 0 5.905 1H3.732a1 1 0 0 0-.866 1.5L6.04 8Zm4.889 0L5.69 13.5H4.598L7.34 8.75L7.773 8l-.433-.75L4.598 2.5H5.69L10.93 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}