package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotateNinetyDegrees(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1664H0V128zm1920 1536V256H128v1408h1792zM896 755L256 995V892l128-48V564l-128-48V413l640 240v102zm-384 41l247-92l-247-92v184zm1277 417l-317 317l-317-317l90-90l163 163V384h128v902l163-163l90 90z"/>`),
		g.Group(children),
	)
}