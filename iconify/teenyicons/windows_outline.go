package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 3.5l-.105-.489A.5.5 0 0 0 0 3.5h.5Zm14-3h.5a.5.5 0 0 0-.605-.489L14.5.5Zm0 14l-.07.495A.5.5 0 0 0 15 14.5h-.5Zm-14-2H0a.5.5 0 0 0 .43.495L.5 12.5Zm.105-8.511l14-3l-.21-.978l-14 3l.21.978ZM14 .5v14h1V.5h-1Zm.57 13.505l-14-2l-.14.99l14 2l.14-.99ZM1 12.5v-9H0v9h1ZM.5 8h14V7H.5v1ZM6 2v11h1V2H6Z"/>`),
		g.Group(children),
	)
}