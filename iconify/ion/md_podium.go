package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdPodium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M32 224h128v192H32z" fill="currentColor"/><path d="M192 128h128v288H192z" fill="currentColor"/><path d="M352 288h128v128H352z" fill="currentColor"/>`),
		g.Group(children),
	)
}