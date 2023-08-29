package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareFontAwesome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M384.5 32.5h-320c-35.3 0-64 28.7-64 64v320c0 35.3 28.7 64 64 64h320c35.3 0 64-28.7 64-64v-320c0-35.3-28.7-64-64-64zm-48 280c-31.6 11.2-41.2 16-59.8 16c-31.4 0-43.2-16-74.6-16c-10.2 0-18.2 1.6-25.6 4v-32c7.4-2.2 15.4-4 25.6-4c31.2 0 43.2 16 74.6 16c10.2 0 17.8-1.4 27.8-4.6v-96c-10 3.2-17.6 4.6-27.8 4.6c-31.4 0-43.2-16-74.6-16c-25.4 0-37.4 10.4-57.6 14.4v153.6c0 8.8-7.2 16-16 16s-16-7.2-16-16v-192c0-8.8 7.2-16 16-16s16 7.2 16 16v6.4c20.2-4 32.2-14.4 57.6-14.4c31.2 0 43.2 16 74.6 16c18.6 0 28.2-4.8 59.8-16v160z"/>`),
		g.Group(children),
	)
}