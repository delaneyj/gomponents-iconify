package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#8d1b3d" d="M0 0h640v480H0z"/><path fill="#fff" d="M0 0v480h158.4l97.8-26.7l-97.8-26.6l97.7-26.7l-97.7-26.7l97.7-26.6l-97.7-26.7l97.8-26.7l-97.8-26.6l97.7-26.7l-97.7-26.7l97.7-26.6l-97.7-26.7l97.8-26.7l-97.8-26.6L256.1 80l-97.7-26.7l97.8-26.6L158.3 0H0z"/>`),
		g.Group(children),
	)
}