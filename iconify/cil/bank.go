package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M247.759 14.358L16 125.946V184h480v-58.362ZM464 152H48v-5.946l200.241-96.412L464 146.362ZM48 408h416v32H48zm-32 56h480v32H16zm40-248h32v160H56zm368 0h32v160h-32zm-96 0h32v160h-32zm-176 0h32v160h-32zm88 0h32v160h-32z"/>`),
		g.Group(children),
	)
}