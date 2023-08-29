package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M280 81.5V48h-48v33.5a191 191 0 0 0-84.43 32.13L120 86l-34 34l25.59 25.59A191.17 191.17 0 0 0 64 272c0 105.87 86.13 192 192 192s192-86.13 192-192c0-97.74-73.42-178.66-168-190.5ZM256 320a48 48 0 0 1-16-93.25V136h32v90.75A48 48 0 0 1 256 320Z"/>`),
		g.Group(children),
	)
}