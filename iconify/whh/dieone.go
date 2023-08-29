package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dieone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 959H128q-53 0-90.5-37.5T0 831V127q0-53 37.5-90T128 0h704q53 0 90.5 37.5T960 127v704q0 53-37.5 90.5T832 959zM480 383q-40 0-68 28.5t-28 68t28 67.5t68 28t68-28t28-67.5t-28-68t-68-28.5z"/>`),
		g.Group(children),
	)
}