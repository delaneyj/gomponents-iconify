package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCuba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M10.8 26h50.6c-.9-4.4-2.8-8.5-5.4-12H10.8v12m0 12v12H56c2.6-3.5 4.5-7.6 5.4-12H10.8z"/><path fill="#428bc1" d="M61.4 26H10.8v12h50.6c.4-1.9.6-3.9.6-6s-.2-4.1-.6-6M10.8 14H56C50.5 6.7 41.8 2 32 2c-8.3 0-15.8 3.4-21.2 8.8V14m0 36v3.2C16.2 58.6 23.7 62 32 62c9.8 0 18.5-4.7 24-12H10.8z"/><path fill="#ed4c5c" d="M10.8 10.8C5.4 16.2 2 23.7 2 32s3.4 15.8 8.8 21.2L32 32L10.8 10.8z"/><path fill="#f9f9f9" d="m10 38l4-2.8l4 2.8l-1.5-4.6l4-2.9h-5L14 26l-1.5 4.5h-5l4 2.9z"/>`),
		g.Group(children),
	)
}