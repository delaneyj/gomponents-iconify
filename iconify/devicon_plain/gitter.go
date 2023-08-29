package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M96.8 25.6H107v51.2H96.8zm-25.6 0h10.2V128H71.2zm-25.6 0h10.2V128H45.6zM20 0h10.2v76.8H20z"/>`),
		g.Group(children),
	)
}