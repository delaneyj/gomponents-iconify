package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M24 10L12 18L24 26L36 18L24 10Z"/><path d="M24 10L36 18L41 11L30 4L24 10Z"/><path d="M24 10L12 18L7 11L18 4L24 10Z"/><path d="M43 22L36 18L24 26L31 31L43 22Z"/><path d="M5 22L12 18L24 26L17 31L5 22Z"/><path stroke-linecap="round" d="M36 28V37L24 44L12 37V28"/></g>`),
		g.Group(children),
	)
}