package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 614V94h680v520H0zm59-343l121-47v-27L59 149v30l86 31l-86 32v29zm132 35h119v-30H191v30z"/>`),
		g.Group(children),
	)
}