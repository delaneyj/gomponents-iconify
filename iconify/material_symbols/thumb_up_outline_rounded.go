package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21V8l6-5.95q.375-.375.888-.437t.987.187q.475.25.7.7t.1.925L14.55 8H21q.8 0 1.4.6T23 10v2q0 .175-.05.375t-.1.375l-3 7.05q-.225.5-.75.85T18 21H7ZM9 8.85V19h9l3-7v-2h-9l1.35-5.5L9 8.85ZM4 21q-.825 0-1.413-.588T2 19v-9q0-.825.588-1.413T4 8h3v2H4v9h3v2H4Zm5-2V8.85V19Z"/>`),
		g.Group(children),
	)
}