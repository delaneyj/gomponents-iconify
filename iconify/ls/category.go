package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Category(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M83 0L28 56h496L469 0H83zm468 276H0V83h551v193zm-165-55v-55H166v55h27v-28h165v28h28zm165 276H0V304h551v193zm-165-55v-56H166v56h27v-28h165v28h28zm165 275H0V525h551v192zm-165-55v-55H166v55h27v-28h165v28h28z"/>`),
		g.Group(children),
	)
}