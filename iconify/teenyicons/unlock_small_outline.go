package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m9.121 4.121l.354-.353l-.354.353ZM5.5 7h4V6h-4v1Zm4.5.5v3h1v-3h-1ZM9.5 11h-4v1h4v-1ZM5 10.5v-3H4v3h1Zm.5.5a.5.5 0 0 1-.5-.5H4A1.5 1.5 0 0 0 5.5 12v-1Zm4.5-.5a.5.5 0 0 1-.5.5v1a1.5 1.5 0 0 0 1.5-1.5h-1ZM9.5 7a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 9.5 6v1Zm-4-1A1.5 1.5 0 0 0 4 7.5h1a.5.5 0 0 1 .5-.5V6Zm.5.5v-.879H5V6.5h1Zm2.768-2.025l.378.379l.708-.708l-.38-.378l-.706.707ZM7.62 4c.43 0 .843.17 1.147.475l.707-.707A2.621 2.621 0 0 0 7.62 3v1ZM6 5.621C6 4.726 6.726 4 7.621 4V3A2.621 2.621 0 0 0 5 5.621h1Z"/>`),
		g.Group(children),
	)
}