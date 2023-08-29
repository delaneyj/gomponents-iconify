package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxUnreadLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M14 2.005C13.385 2 12.72 2 12 2C7.286 2 4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12c0-.72 0-1.385-.005-2"/><circle cx="19" cy="5" r="3"/><path stroke-linecap="round" d="M2 13h3.16c.905 0 1.358 0 1.756.183c.398.183.692.527 1.281 1.214l.606.706c.589.687.883 1.031 1.281 1.214c.398.183.85.183 1.756.183h.32c.905 0 1.358 0 1.756-.183c.398-.183.692-.527 1.281-1.214l.606-.706c.589-.687.883-1.031 1.281-1.214c.398-.183.85-.183 1.756-.183H22" opacity=".5"/></g>`),
		g.Group(children),
	)
}