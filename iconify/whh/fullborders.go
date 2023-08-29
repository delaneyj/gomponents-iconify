package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullborders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M0 960V0h960v960H0zm64-64h384V512H64v384zm0-832v384h384V64H64zm832 0H512v384h384V64zm0 448H512v384h384V512z"/>`),
		g.Group(children),
	)
}