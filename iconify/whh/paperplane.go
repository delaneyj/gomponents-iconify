package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperplane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 512v352q0 13-9.5 22.5T864 896h-64L656 752l-272 272L0 0l1024 384zm-64 0L128 128l704 704V512z"/>`),
		g.Group(children),
	)
}