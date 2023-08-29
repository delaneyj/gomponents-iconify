package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PigNose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M10 25.5c-1.52 0-2.75-1.23-2.75-2.75v-3.88c0-1.52 1.23-2.75 2.75-2.75s2.75 1.23 2.75 2.75v3.88c0 1.52-1.23 2.75-2.75 2.75Zm9.25-2.75c0 1.52 1.23 2.75 2.75 2.75s2.75-1.23 2.75-2.75v-3.88c0-1.52-1.23-2.75-2.75-2.75s-2.75 1.23-2.75 2.75v3.88Z"/><path d="M1 21.776C1.12 13.514 7.894 8 16 8c8.107 0 14.88 5.514 15 13.776v.223A8.998 8.998 0 0 1 22 31H10c-4.972 0-9-4.028-9-9v-.224Zm2 .021V22c0 3.867 3.132 7 7 7h12c3.868 0 7-3.132 7-7v-.202C28.895 14.824 23.211 10 16 10S3.104 14.823 3 21.797Z"/></g>`),
		g.Group(children),
	)
}