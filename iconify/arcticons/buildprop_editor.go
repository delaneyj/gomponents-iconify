package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildpropEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.736 28.315a3.244 3.244 0 0 1-3.211-3.256v-2.117a3.212 3.212 0 1 1 6.423 0v2.117a3.244 3.244 0 0 1-3.212 3.256Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.043 22.944a3.244 3.244 0 0 1 3.211-3.256m-3.211 0v8.629m-9.212-3.257a3.212 3.212 0 1 0 6.423 0v-2.116a3.212 3.212 0 1 0-6.423 0m0-3.256v13.025m23.915-7.653a3.212 3.212 0 1 0 6.423 0v-2.116a3.212 3.212 0 1 0-6.423 0m-.001-3.256v13.025" class="b"/><ellipse cx="7.446" cy="32.459" fill="currentColor" rx=".74" ry=".75"/>`),
		g.Group(children),
	)
}