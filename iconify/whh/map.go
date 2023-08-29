package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 1024V640l320-128v384zm0-896L1024 0v448L704 576V128zM384 512l256 128v384L384 896V512zm0-512l256 128v448L384 448V0zM0 640l320-128v384L0 1024V640zm0-512L320 0v448L0 576V128z"/>`),
		g.Group(children),
	)
}