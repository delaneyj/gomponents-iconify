package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Allcast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="25.227" x="5.5" y="9.447" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.682"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.102 38.553h17.796M5.5 30.675a4 4 0 0 1 4 3.999h0m-4-6.999a6.998 6.998 0 0 1 6.999 6.999h0M5.5 18.677a15.997 15.997 0 0 1 15.997 15.997"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24.676a9.998 9.998 0 0 1 9.998 9.998h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 21.677a12.997 12.997 0 0 1 12.997 12.997"/>`),
		g.Group(children),
	)
}