package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batterycharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 512h-96v128q0 53-37.5 90.5T768 768H128q-53 0-90.5-37.5T0 640V128q0-53 37.5-90.5T128 0h640q53 0 90.5 37.5T896 128v128h96q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5T992 512zM768 186q0-24-17-41t-42-17H187q-25 0-42 17t-17 41v396q0 24 17 41t42 17h522q25 0 42-17t17-41V186zM384 384H192l320-128v128h192L384 512V384z"/>`),
		g.Group(children),
	)
}