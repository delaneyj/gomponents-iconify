package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoClicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.675 31.338l8.579-3.41a3.514 3.514 0 0 0-2.473-6.579l-1.07.402l5.259-10.4a3.514 3.514 0 1 0-6.272-3.17L29.715 12.1a3.49 3.49 0 0 0-2.849-.12a3.48 3.48 0 0 0-1.277-.94a3.476 3.476 0 0 0-1.507-.263a3.498 3.498 0 0 0-1.14-.794a3.51 3.51 0 0 0-4.624 1.822l-5.243 12.062"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.2 14.853A5.268 5.268 0 0 0 34.833 4.5a5.266 5.266 0 0 0-4.907 7.18M15.163 23.965l12.512 7.373a2.249 2.249 0 0 1 .93 2.812l-3.37 7.976a2.249 2.249 0 0 1-3.184 1.079l-13.014-7.41a2.249 2.249 0 0 1-.909-2.94L12 24.916a2.249 2.249 0 0 1 3.163-.951Z"/>`),
		g.Group(children),
	)
}