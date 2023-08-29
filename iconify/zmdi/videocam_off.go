package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M405 99v228L167 88h132q8 0 14.5 6.5T320 109v75zM27 3l378 378l-27 27l-68-68q-6 4-11 4H43q-9 0-15.5-6.5T21 323V109q0-8 6.5-14.5T43 88h15L0 30z"/>`),
		g.Group(children),
	)
}