package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenMailboxWithLoweredFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#b0bdc6" d="m38 64l-8.8-1.2v-49H38z"/><path fill="#7d8b91" d="M44.3 13.8H38V64l6.3-6.1z"/><path fill="#333" d="M46.8 0c-8.2 0-35.6 9.5-35.6 9.5l24.6 40.4l28.3-17V17C64 6.3 57.6 0 46.8 0z"/><path fill="#697277" d="M17.2 8.5C6.4 8.5 0 16 0 26.8v17l35.7 6.1V29.2c0-10.8-7.6-20.7-18.5-20.7z"/><path fill="#94989b" d="M35.7 49.9L0 43.8l35.7-16.4z"/><path fill="#ed4c5c" d="m62.7 14.2l-26 12.4v3.3l19.2-9.6v9.9l6.8-3.8z"/>`),
		g.Group(children),
	)
}