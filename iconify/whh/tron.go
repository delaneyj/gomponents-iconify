package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.232 256h-192q-53 0-90.5 37.5t-37.5 90.5v512q0 26-18.5 45t-45.5 19h-128q-27 0-45.5-19t-18.5-45V384q0-104 51.5-192.5t140-140T768.232 0h192q26 0 45 19t19 45v128q0 27-18.5 45.5t-45.5 18.5zm0-192h-192q-85 0-159 44t-117.5 117.5t-43.5 158.5v512h128V384q0-73 60-132.5t132-59.5h192V64zm-640 192h-256q-26 0-45-19t-19-45V64q0-27 18.5-45.5T64.232 0h256q27 0 45.5 19t18.5 45v128q0 27-18.5 45.5t-45.5 18.5zm0-192h-256v128h256V64z"/>`),
		g.Group(children),
	)
}