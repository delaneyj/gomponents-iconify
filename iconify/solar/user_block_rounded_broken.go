package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockRoundedBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="6" r="4"/><path stroke-linecap="round" d="m19.95 16.05l-3.9 3.9"/><circle cx="18" cy="18" r="3"/><path stroke-linecap="round" d="M13 20.96c-.327.026-.66.04-1 .04c-3.866 0-7-1.79-7-4c0-.345.077-.68.22-1m9.28-2.737c-.776-.17-1.62-.263-2.5-.263c-1.074 0-2.09.138-3 .385"/></g>`),
		g.Group(children),
	)
}