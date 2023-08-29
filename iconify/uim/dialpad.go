package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dialpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="9" r="2" fill="currentColor" opacity=".5"/><circle cx="12" cy="3" r="2" fill="currentColor" opacity=".5"/><circle cx="12" cy="15" r="2" fill="currentColor" opacity=".5"/><circle cx="6" cy="9" r="2" fill="currentColor" opacity=".5"/><circle cx="6" cy="3" r="2" fill="currentColor" opacity=".5"/><circle cx="6" cy="15" r="2" fill="currentColor" opacity=".5"/><circle cx="18" cy="9" r="2" fill="currentColor" opacity=".5"/><circle cx="18" cy="3" r="2" fill="currentColor" opacity=".5"/><circle cx="18" cy="15" r="2" fill="currentColor" opacity=".5"/><circle cx="12" cy="21" r="2" fill="currentColor" opacity=".5"/>`),
		g.Group(children),
	)
}