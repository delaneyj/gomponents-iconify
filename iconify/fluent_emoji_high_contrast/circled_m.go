package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M23.238 10.814a1.375 1.375 0 0 0-2.27-1.044l-4.918 4.216a.125.125 0 0 1-.16.002L11 10.048a1.375 1.375 0 0 0-2.237 1.072v10.454a1.375 1.375 0 1 0 2.75 0v-7.32c0-.105.121-.164.203-.098l3.422 2.756a1.375 1.375 0 0 0 1.757-.027l3.388-2.905a.125.125 0 0 1 .206.095v7.5a1.375 1.375 0 1 0 2.75 0V10.813Z"/><path d="M16 1C7.716 1 1 7.716 1 16c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15ZM3 16C3 8.82 8.82 3 16 3s13 5.82 13 13s-5.82 13-13 13S3 23.18 3 16Z"/></g>`),
		g.Group(children),
	)
}