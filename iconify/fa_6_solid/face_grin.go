package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceGrin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512zm132.1-199.2c12.3-3.8 24.3 6.9 19.3 18.7c-25 59.1-83.2 100.5-151.1 100.5s-126.2-41.4-151.1-100.5c-5-11.8 7-22.5 19.3-18.7c39.7 12.2 84.5 19 131.8 19s92.1-6.8 131.8-19zM144.4 208a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}