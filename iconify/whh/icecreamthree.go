package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icecreamthree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 576q0 27-18.5 45.5T672 640h-11l-53 320q0 27-18.5 45.5T544 1024H224q-27 0-45.5-18.5T160 960l-53-320H96q-27 0-45.5-18.5T32 576L0 448q0-27 18.5-45.5T64 384h640q27 0 45.5 18.5T768 448zm-64-256H96q-13 0-22.5-9.5T64 288q0-86 43.5-152.5t116-101T384.5 0t161 34.5t115.5 101T704 288q0 13-9.5 22.5T672 320z"/>`),
		g.Group(children),
	)
}