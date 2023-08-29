package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Touchpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 896H512V640h256q27 0 45.5-18.5T832 576V192q0-26-18.5-45T768 128H192q-27 0-45.5 19T128 192v384q0 27 19 45.5t45 18.5h256v256H128q-53 0-90.5-37.5T0 768V128q0-53 37.5-90.5T128 0h704q53 0 90.5 37.5T960 128v640q0 53-37.5 90.5T832 896z"/>`),
		g.Group(children),
	)
}