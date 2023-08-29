package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOffTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1536 760l512-256v1040l-250-125q-18-9-33-15t-31-13t-29-15t-32-20q-15-10-26-20t-20-21t-19-22t-23-25l-177-177V640H957L829 512h707v248zm384 577V711l-384 193v240l384 193zM19 109l90-90l1920 1920l-90 90l-494-493H0V512h421L19 109zm109 1299h1189L549 640H128v768z"/>`),
		g.Group(children),
	)
}