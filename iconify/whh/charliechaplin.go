package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Charliechaplin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 512q0 35-68.5 64.5T769 623t-257 17t-257-17t-186.5-46.5T0 512q0-48 129-85q-1-22-1-27q0-109 51.5-201t140-145.5T512 0t192.5 53.5t140 145.5T896 400q0 5-1 27q129 37 129 85zM576 896l64 128H384l64-128h128z"/>`),
		g.Group(children),
	)
}