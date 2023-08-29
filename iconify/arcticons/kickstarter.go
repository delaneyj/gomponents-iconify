package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kickstarter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.97 24.009l3.547-3.523a9.332 9.332 0 0 0 0-13.24a9.485 9.485 0 0 0-13.33 0l-1.292 1.282a9.435 9.435 0 0 0-17.167 5.33v20.285a9.435 9.435 0 0 0 17.167 5.329l1.291 1.282a9.485 9.485 0 0 0 13.33 0a9.332 9.332 0 0 0 0-13.24L34.97 24.01"/>`),
		g.Group(children),
	)
}