package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Covidcheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.385 11.23h5.004a1.38 1.38 0 0 1 1.38 1.38v5.005m-25.539 0v-5.004a1.38 1.38 0 0 1 1.38-1.38h5.005m0 25.539h-5.004a1.38 1.38 0 0 1-1.38-1.38v-5.006M19.44 24l3.374 3.648l7.114-7.297m2.243 16.418h-1.786m8.451 3.961v-7.155m0 11.925c7.065-2.715 6.66-8.154 6.66-11.4v-2.46a25.214 25.214 0 0 0-6.66-1.14a25.214 25.214 0 0 0-6.66 1.14v2.46c0 3.246-.405 8.685 6.66 11.4Zm-3.578-8.347h7.155"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.54 42.743a21.503 21.503 0 1 1 10.959-18.97m.001.173a21.46 21.46 0 0 1-1.265 7.338"/>`),
		g.Group(children),
	)
}