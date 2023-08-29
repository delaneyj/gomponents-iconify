package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Personio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.18 42.502h29.579M26.717 12.444l-6.454 24.083m-3.906-9.256c6.428.16 35.593-19.663 22.253-21.668C31.698 4.24 7.135 17.112 7.587 22.849c.052.667.8.761 1.866 1.114"/><circle cx="30.733" cy="35.307" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}