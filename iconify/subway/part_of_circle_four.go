package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartOfCircleFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0v256H0c0 141.4 114.6 256 256 256s256-114.6 256-256S397.4 0 256 0z"/>`),
		g.Group(children),
	)
}