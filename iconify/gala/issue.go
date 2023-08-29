package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Issue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaIssue0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><circle id="galaIssue1" cx="128" cy="128" r="40"/><circle id="galaIssue2" cx="128" cy="128" r="112"/></g>`),
		g.Group(children),
	)
}