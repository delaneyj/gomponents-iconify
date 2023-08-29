package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M37.972 27.975A12.931 12.931 0 0 0 40 21c0-7.18-5.82-13-13-13c-2.567 0-4.96.744-6.975 2.027a16.953 16.953 0 0 0-3.954.698a15.069 15.069 0 0 1 3.326-2.658A14.93 14.93 0 0 1 26.95 6H27c8.284 0 15 6.716 15 15c0 .135-.002.269-.005.403l-.002.048a14.921 14.921 0 0 1-2.06 7.152a15.065 15.065 0 0 1-2.658 3.326c.381-1.263.62-2.587.697-3.954Z"/><path fill-rule="evenodd" d="M25.772 22.667A4.001 4.001 0 0 0 22 20v-1h-2v1a4 4 0 0 0 0 8v4c-.87 0-1.611-.555-1.887-1.333a1 1 0 1 0-1.885.666A4.001 4.001 0 0 0 20 34v1h2v-1a4 4 0 0 0 0-8v-4c.87 0 1.611.555 1.887 1.333a1 1 0 1 0 1.885-.666ZM20 22a2 2 0 1 0 0 4v-4Zm2 10a2 2 0 1 0 0-4v4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M36 27c0 8.284-6.716 15-15 15c-8.284 0-15-6.716-15-15c0-8.284 6.716-15 15-15c8.284 0 15 6.716 15 15Zm-2 0c0 7.18-5.82 13-13 13S8 34.18 8 27s5.82-13 13-13s13 5.82 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}