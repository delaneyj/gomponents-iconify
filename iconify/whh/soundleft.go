package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soundleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 767H800q-13 0-22.5-9t-9.5-23V287q0-13 9.5-22.5T800 255t22.5 9.5T832 287v416h160q13 0 22.5 9.5t9.5 23t-9.5 22.5t-22.5 9zm-427 244L324 768h-4V255h4L565 12q30-30 75 16v968q-45 46-75 15zM0 704V319q0-26 18.5-45T64 255h192v513H64q-27 0-45.5-19T0 704z"/>`),
		g.Group(children),
	)
}