package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Embed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m467 595l90 90l-338 339l338 339l-90 90l-430-429l430-429zm1114 0l430 429l-430 429l-90-90l338-339l-338-339l90-90zM701 1792l512-1536h134L835 1792H701z"/>`),
		g.Group(children),
	)
}