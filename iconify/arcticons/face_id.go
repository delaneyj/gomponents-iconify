package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceId(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14 5.5H7.5a2 2 0 0 0-2 2V14m37 0V7.5a2 2 0 0 0-2-2H34m0 37h6.5a2 2 0 0 0 2-2V34m-37 0v6.5a2 2 0 0 0 2 2H14m2.281-22.437V16.5m15.438 3.563V16.5M21.625 28.375h1.188A2.375 2.375 0 0 0 25.188 26v-9.5m5.937 15.773C29.189 34.839 26.233 35.5 24 35.5c-2.233 0-5.189-.66-7.125-3.227"/>`),
		g.Group(children),
	)
}