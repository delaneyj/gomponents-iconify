package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubnetAclRules(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 14h12v2H18zm0 5h8v2h-8zm0-10h12v2H18z"/><path fill="currentColor" d="M22 24v4H6V16h8v-2h-4V8a4 4 0 0 1 7.668-1.6l1.832-.8A6.001 6.001 0 0 0 8 8v6H6a2.002 2.002 0 0 0-2 2v12a2.002 2.002 0 0 0 2 2h16a2.002 2.002 0 0 0 2-2v-4Z"/>`),
		g.Group(children),
	)
}