package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypeSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.129 3.774c0-.427.347-.774.774-.774h6.194a.774.774 0 0 1 0 1.548H8.774v7.484a.774.774 0 0 1-1.548 0V4.548H4.903a.774.774 0 0 1-.774-.774Zm-.743 1.517a.774.774 0 0 1 0 1.095L1.87 7.903L3.386 9.42a.774.774 0 0 1-1.095 1.095L.227 8.451a.774.774 0 0 1 0-1.095L2.29 5.29a.774.774 0 0 1 1.095 0Zm9.228 0a.774.774 0 0 1 1.095 0l2.064 2.065a.774.774 0 0 1 0 1.095l-2.064 2.064a.774.774 0 0 1-1.095-1.095l1.517-1.517l-1.517-1.517a.774.774 0 0 1 0-1.095Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}