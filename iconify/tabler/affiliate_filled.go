package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AffiliateFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18.5 3a2.5 2.5 0 1 1-.912 4.828l-4.556 4.555a5.475 5.475 0 0 1 .936 3.714l2.624.787a2.5 2.5 0 1 1-.575 1.916l-2.623-.788a5.5 5.5 0 0 1-10.39-2.29L3 15.5l.004-.221a5.5 5.5 0 0 1 2.984-4.673L5.2 7.982a2.498 2.498 0 0 1-2.194-2.304L3 5.5l.005-.164a2.5 2.5 0 1 1 4.111 2.071l.787 2.625a5.475 5.475 0 0 1 3.714.936l4.555-4.556a2.487 2.487 0 0 1-.167-.748L16 5.5l.005-.164A2.5 2.5 0 0 1 18.5 3z"/></g>`),
		g.Group(children),
	)
}