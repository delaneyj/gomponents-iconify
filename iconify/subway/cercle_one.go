package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CercleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0zm0 426.7c-94.3 0-170.7-76.4-170.7-170.7S161.7 85.3 256 85.3S426.7 161.7 426.7 256S350.3 426.7 256 426.7zm0-256c-47.1 0-85.3 38.2-85.3 85.3s38.2 85.3 85.3 85.3s85.3-38.2 85.3-85.3s-38.2-85.3-85.3-85.3z"/>`),
		g.Group(children),
	)
}