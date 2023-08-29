package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22s-9-4-9-11V5l9-3l9 3v6c0 7-9 11-9 11Zm0-8a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-6V5m0 12v-3m-6-3h3m6 0h3M8 7l2 2m4 4l2 2m0-8l-2 2m-4 4l-2 2"/>`),
		g.Group(children),
	)
}