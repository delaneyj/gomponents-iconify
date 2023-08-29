package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XPack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18.266 26.068l7.839-7.854l4.469 4.479a4.763 4.763 0 0 1 0 6.734l-1.104 1.104a4.752 4.752 0 0 1-6.734 0zM30.563 2.531l-1.109-1.104a4.763 4.763 0 0 0-6.734 0l-6.719 6.734l-6.734-6.734a4.763 4.763 0 0 0-6.734 0L1.429 2.531a4.763 4.763 0 0 0 0 6.734l6.734 6.734l-6.734 6.734a4.763 4.763 0 0 0 0 6.734l1.104 1.104a4.763 4.763 0 0 0 6.734 0L30.574 9.264a4.763 4.763 0 0 0 0-6.734z"/>`),
		g.Group(children),
	)
}