package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatermelonOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 4L41 33.92C41 33.92 36.0457 38 24 38C11.9543 38 7 33.92 7 33.92L24 4Z"/><circle cx="24" cy="17" r="2" fill="#fff"/><circle cx="27" cy="23" r="2" fill="#fff"/><circle cx="21" cy="23" r="2" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-width="4" d="M41 39.9199C41 39.9199 36.0457 43.9999 24 43.9999C11.9543 43.9999 7 39.9199 7 39.9199"/></g>`),
		g.Group(children),
	)
}