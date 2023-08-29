package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateLocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0 0 18a9.003 9.003 0 0 0 8.252-5.4l1.832.8A11.002 11.002 0 0 1 12 23C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11v2l-3.6-2.7l1.2-1.6l.128.096A9.004 9.004 0 0 0 12 3Zm0 5.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25Zm3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75H7.499V17h9v-6.5H15.25Zm-.752 2V15h-5v-2.5h5Z"/>`),
		g.Group(children),
	)
}