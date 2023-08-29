package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailProhibitedSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 5.5a4.5 4.5 0 1 0 9 0a4.5 4.5 0 0 0-9 0Zm1 0a3.5 3.5 0 0 1 5.596-2.803l-4.9 4.9A3.484 3.484 0 0 1 7 5.5ZM10.5 9c-.786 0-1.512-.26-2.096-.697l4.9-4.9A3.5 3.5 0 0 1 10.5 9Zm1.5 4v-2.207c.349-.099.683-.23 1-.393V13a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2.022a5.57 5.57 0 0 0 0 1H3a1 1 0 0 0-1 1v.74l5 2.692l.544-.293c.344.22.714.402 1.104.542l-1.41.76a.5.5 0 0 1-.475 0L2 8.875V13a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1Z"/>`),
		g.Group(children),
	)
}