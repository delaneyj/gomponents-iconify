package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M4 32.0833C4 30.8812 4.266 29.6884 4.97123 28.7148C7.04541 25.8514 12.6701 20 24 20C35.3299 20 40.9546 25.8514 43.0288 28.7148C43.734 29.6884 44 30.8812 44 32.0833C44 36.4556 40.4556 40 36.0833 40H11.9167C7.54441 40 4 36.4556 4 32.0833Z"/><path stroke="#000" d="M12 9L12 13"/><path stroke="#fff" d="M14 22L14 26"/><path stroke="#000" d="M36 9L36 13"/><path stroke="#fff" d="M34 22L34 26"/><path stroke="#000" d="M24 7L24 13"/><path stroke="#fff" d="M24 20L24 28"/><path stroke="#000" d="M40 25.4434C36.9058 22.7787 31.8075 20 24 20C16.1925 20 11.0942 22.7787 8 25.4434"/></g>`),
		g.Group(children),
	)
}