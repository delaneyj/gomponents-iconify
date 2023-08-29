package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EsPvOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#d52b1e" d="M0 0h512v512H0z"/><path fill="#009b48" d="M0 0h42.5l106.7 106.7L256 213.4L469.5 0H512v42.5L405.3 149.2L298.6 256L512 469.5V512h-42.5L362.8 405.3L256 298.6L42.5 512H0v-42.5l106.7-106.7L213.4 256L0 42.5V21.3z"/><path fill="#fff" d="M221.9 0h68.2v221.9H512v68.2H290.1V512h-68.2V290.1H0v-68.2h221.9v-111z"/>`),
		g.Group(children),
	)
}