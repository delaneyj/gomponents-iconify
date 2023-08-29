package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shareme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.4 17.33a9.43 9.43 0 1 1 6.67 16.1c-5.21 0-7.42-4.95-10.07-9.43c-2.83-4.78-4.86-9.43-10.07-9.43a9.43 9.43 0 1 0 6.6 16.16"/>`),
		g.Group(children),
	)
}