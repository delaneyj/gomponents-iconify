package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windowseight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M478 922V546h546v478zm0-819L1024 0v478H478V103zm-68 375H0V205l410-90v363zm0 431L0 819V546h410v363z"/>`),
		g.Group(children),
	)
}