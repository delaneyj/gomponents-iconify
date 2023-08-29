package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElasticStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M.094 3.266A3.202 3.202 0 0 1 3.266.094h25.469a3.202 3.202 0 0 1 3.172 3.172v6.375H.094zm0 19.093h15.104v9.547H3.266a3.172 3.172 0 0 1-3.172-3.094zm16.708 0h15.104v6.375c0 1.703-1.469 3.172-3.172 3.172H16.802zM.094 11.229h31.813v9.542H.094z"/>`),
		g.Group(children),
	)
}