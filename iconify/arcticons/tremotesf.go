package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tremotesf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.687 7.709H9.09a2.485 2.485 0 0 0-2.209 2.209v4.97a2.485 2.485 0 0 0 2.209 2.209h11.597m6.626 0H38.91a2.485 2.485 0 0 0 2.209-2.209v-4.97A2.485 2.485 0 0 0 38.91 7.71H27.313"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.687 5.5h6.626v23.746h4.97L24 36.978l-8.284-7.732h4.97Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.687 22.62h-9.94a2.485 2.485 0 0 0-2.21 2.208v15.463a2.485 2.485 0 0 0 2.21 2.209h26.507a2.485 2.485 0 0 0 2.209-2.209V24.828a2.485 2.485 0 0 0-2.21-2.209h-9.94"/>`),
		g.Group(children),
	)
}