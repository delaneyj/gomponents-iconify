package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Habitnow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.372 29.858a2.94 2.94 0 0 1-4.166 0l-5.43-5.43a1.662 1.662 0 0 0-2.381 0l-.893.892c-.67.67-.67 1.711 0 2.381l5.505 5.505a2.829 2.829 0 0 0 4.017 0l15.474-15.474c.67-.67.67-1.71 0-2.38l-.893-.893a1.662 1.662 0 0 0-2.38 0l-14.21 14.21"/>`),
		g.Group(children),
	)
}