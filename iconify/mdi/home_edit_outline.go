package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeEditOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.041 11.14c.142 0 .276.057.38.163l1.276 1.276a.541.541 0 0 1 0 .77l-1 1l-2.046-2.046l1-1a.567.567 0 0 1 .39-.163Zm-1.978 1.74l2.046 2.052L15.063 21H13v-2.063l6.063-6.057ZM12 5.688l-5 4.5V18h4v2H5v-8H2l10-9l7.459 6.713L17 12.172v-1.984l-5-4.5Z"/>`),
		g.Group(children),
	)
}