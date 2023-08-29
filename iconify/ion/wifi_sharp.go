package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="42" d="M332.69 320a115 115 0 0 0-152.8 0m213.85-61a201.26 201.26 0 0 0-274.92 0"/><path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="42" d="M448 191.52a288 288 0 0 0-383.44 0"/><path fill="currentColor" d="M300.67 384L256 433l-44.34-49a56.73 56.73 0 0 1 88.92 0Z"/>`),
		g.Group(children),
	)
}