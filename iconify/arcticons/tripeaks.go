package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tripeaks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.508 27.107l14.534-13.425l6.021 5.991m-5.235 6.157l7.264-8.477L39.5 27.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.5 33.157l11.58-10.911l12.213 12.072"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.299 5.5H40.7a1.799 1.799 0 0 1 1.8 1.799V40.7a1.799 1.799 0 0 1-1.799 1.799H7.3a1.799 1.799 0 0 1-1.8-1.798V7.3a1.799 1.799 0 0 1 1.799-1.8Z"/>`),
		g.Group(children),
	)
}