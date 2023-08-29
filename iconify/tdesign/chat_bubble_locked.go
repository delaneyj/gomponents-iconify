package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatBubbleLocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.959 10.959 0 0 1 1 12Zm11-3.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25Zm3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75H7.499V17h9v-6.5H15.25Zm-.752 2V15h-5v-2.5h5Z"/>`),
		g.Group(children),
	)
}