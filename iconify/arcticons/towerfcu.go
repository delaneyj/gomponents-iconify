package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Towerfcu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.46 4.5h29.08M9.46 10.398h11.526m0 6.019V4.5m6.018 5.898v31.758M9.46 16.417v13c-.01 7.041 4.584 9.66 10.151 12.137l4.374 1.946l4.383-1.946c5.568-2.468 10.172-5.086 10.172-12.147v-12.9l-5.768-.02v22.84"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.46 16.417h11.526v25.739m-5.768-25.739v22.92m11.786-28.939H38.54"/>`),
		g.Group(children),
	)
}