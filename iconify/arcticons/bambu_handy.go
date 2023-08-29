package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BambuHandy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.1 19.01V43.5h14V24.6zm0-14.51v12.141l14 5.59V4.5zM8.9 28.99V4.5h14v18.9zm0 14.51V31.359l14-5.59V43.5z"/>`),
		g.Group(children),
	)
}