package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildkiteIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#30F2A2" d="m0 0l85.333 41.813v85.334L0 85.333zm171.52 0L256 41.813l-84.48 43.52z"/><path fill="#14CC80" d="M171.52 0L85.333 41.813v85.334l86.187-41.814zM256 41.813l-84.48 43.52v85.334l84.48-43.52z"/>`),
		g.Group(children),
	)
}