package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 22c4.97 0 9-2.239 9-5s-4.03-5-9-5s-9 2.239-9 5s4.03 5 9 5Zm0-4.5c4.97 0 9-2.239 9-5s-4.03-5-9-5s-9 2.239-9 5s4.03 5 9 5Zm0-5.5c4.97 0 9-2.239 9-5s-4.03-5-9-5s-9 2.239-9 5s4.03 5 9 5Z"/>`),
		g.Group(children),
	)
}