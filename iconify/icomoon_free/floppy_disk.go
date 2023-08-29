package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 0H0v16h16V2l-2-2zM8 2h2v4H8V2zm6 12H2V2h1v5h9V2h1.172l.828.828V14z"/>`),
		g.Group(children),
	)
}