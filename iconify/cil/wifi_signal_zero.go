package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiSignalZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M237.778 480h36.444L504 151.842V124.4l-.215-.15a432.019 432.019 0 0 0-495.57 0L8 124.4v27.438ZM256 78.128a397.867 397.867 0 0 1 216.144 63.419L256 450.232L39.856 141.547A397.867 397.867 0 0 1 256 78.128Z"/>`),
		g.Group(children),
	)
}