package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trexrunner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.58 13.78V6.36h-1.86V4.5H25.86v1.86H24V17.5h-1.86v1.86h-1.86v1.86h-3.7v1.86h-1.86v1.84H11v-1.84H9.14v-1.86H7.28V17.5H5.42v11.14h1.86v1.86h1.86v1.86H11v1.86h1.86v1.86h1.86v7.42h3.7v-1.86h-1.84v-1.86h1.84v-1.86h1.86v-1.84h1.86v1.84H24v5.58h3.72v-1.86h-1.86v-7.42h1.86v-1.86h1.86v-3.72h1.84v-5.56h1.86v1.84h1.86v-3.7h-3.72V17.5h7.44v-1.86h-5.58v-1.86Zm-13-5.56h-1.86V6.36h1.86Z"/>`),
		g.Group(children),
	)
}