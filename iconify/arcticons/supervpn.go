package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Supervpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.693 0 15.274-7.782 15.274-16.966V6.504C35.267 6.503 24 4.5 24 4.5S12.723 6.503 8.726 6.503v20.03C8.726 35.719 22.307 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.705 27.637c1.043 1.357 2.35 1.863 4.169 1.863h2.516a4.24 4.24 0 0 0 4.241-4.24v-.02a4.24 4.24 0 0 0-4.24-4.24h-2.777a4.245 4.245 0 0 1-4.245-4.245h0a4.255 4.255 0 0 1 4.254-4.255h2.503c1.82 0 3.126.506 4.169 1.863"/>`),
		g.Group(children),
	)
}