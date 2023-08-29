package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 366v530q0 53-37.5 90.5T832 1024H192q-53 0-90.5-37.5T64 896V366q-29-17-46.5-46T0 256q0-70 68.5-128.5t186.5-93T512 0t257 34.5t186.5 93T1024 256q0 35-17.5 64T960 366z"/>`),
		g.Group(children),
	)
}