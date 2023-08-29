package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarezero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-640q0-53-37.5-90.5t-90.5-37.5h-128q-53 0-90.5 37.5t-37.5 90.5v256q0 53 37.5 90.5t90.5 37.5h128q53 0 90.5-37.5t37.5-90.5V384zm-256-64h128q26 0 45 19t19 45v234l-212-294q11-4 20-4zm128 384h-128q-27 0-45.5-18.5t-18.5-45.5V384q0-2 1-9l228 317q-17 12-37 12z"/>`),
		g.Group(children),
	)
}