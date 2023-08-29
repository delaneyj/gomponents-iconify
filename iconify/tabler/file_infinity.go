package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileInfinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.536 17.586a2.123 2.123 0 0 0-2.929 0a1.951 1.951 0 0 0 0 2.828c.809.781 2.12.781 2.929 0c.809-.781-.805.778 0 0l1.46-1.41l1.46-1.419"/><path d="m15.54 17.582l1.46 1.42l1.46 1.41c.809.78-.805-.779 0 0s2.12.781 2.929 0a1.951 1.951 0 0 0 0-2.828a2.123 2.123 0 0 0-2.929 0M14 3v4a1 1 0 0 0 1 1h4"/><path d="M9.5 21H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7l5 5v6"/></g>`),
		g.Group(children),
	)
}