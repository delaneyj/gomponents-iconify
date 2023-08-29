package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LivresDeProches(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 41.676c-1.47-3.758-8.439-10.139-18.5-9.317V6.396c10.061-.822 17.03 5.559 18.5 9.317m0 25.963c1.47-3.758 8.439-10.139 18.5-9.317V6.396c-10.061-.822-17.03 5.559-18.5 9.317m0 0v25.963"/>`),
		g.Group(children),
	)
}