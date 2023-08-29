package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2h8a3.99 3.99 0 0 1 3 1.354A3.99 3.99 0 0 1 15 2h8v19h-1c-2.944 0-5.14.245-6.586.486c-.723.12-1.26.24-1.609.328a10.53 10.53 0 0 0-.472.13l-.017.005h-.002l-.152.051h-2.324l-.152-.05l-.001-.001h-.001l-.017-.006a10.53 10.53 0 0 0-.472-.13a20.76 20.76 0 0 0-1.61-.327C7.14 21.246 4.946 21 2 21H1V2Zm2 2v15.01c2.563.047 4.535.274 5.914.504c.777.13 1.366.26 1.766.36c.125.03.232.06.32.084V6a2 2 0 0 0-2-2H3Zm10 2v13.957a22.771 22.771 0 0 1 2.086-.444c1.379-.23 3.35-.456 5.914-.504V4h-6a2 2 0 0 0-2 2Zm2 2h4v2h-4V8Zm0 3h4v2h-4v-2Z"/>`),
		g.Group(children),
	)
}