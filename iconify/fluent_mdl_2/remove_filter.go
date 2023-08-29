package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 128h1792v1792H128V128zm1664 1664V256H256v1536h1536zM621 1517l-90-90l402-403l-402-403l90-90l403 402l403-402l90 90l-402 403l402 403l-90 90l-403-402l-403 402z"/>`),
		g.Group(children),
	)
}