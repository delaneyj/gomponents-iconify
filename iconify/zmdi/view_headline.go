package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewHeadline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 256v-43h341v43H0zm0 85v-42h341v42H0zm0-170v-43h341v43H0zM0 43h341v42H0V43z"/>`),
		g.Group(children),
	)
}