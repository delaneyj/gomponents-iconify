package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PixelTips(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.987 17.487a12.987 12.987 0 1 0-20.163 10.82v11.618a1.391 1.391 0 0 0 1.391 1.391h11.57a1.391 1.391 0 0 0 1.391-1.391V28.308a12.97 12.97 0 0 0 5.811-10.821Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.824 34.465a1.391 1.391 0 0 0 1.391 1.391h11.57a1.391 1.391 0 0 0 1.391-1.391M17.994 23.688h12.012M24 35.856V23.688m-3.744 17.628v.793a1.391 1.391 0 0 0 1.391 1.391h4.706a1.391 1.391 0 0 0 1.391-1.391v-.793"/>`),
		g.Group(children),
	)
}