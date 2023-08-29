package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Googledrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m376 13l8-13h256l346 576H701zM143 845L0 640L335 81l123 213zm199-205h682L846 896H194z"/>`),
		g.Group(children),
	)
}