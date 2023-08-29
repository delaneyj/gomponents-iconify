package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 700v196l-64 128l-64-128V700q-107-14-195-67T50.5 497.5T0 320q0-56 20.5-116T86 87T192 0Q64 157 64 240q0 124 91.5 217.5T384 571V448l64-128l64 128v123q137-20 228.5-113.5T832 240q0-83-128-240q61 30 106 87t65.5 117T896 320q0 95-50.5 177.5T707 633t-195 67z"/>`),
		g.Group(children),
	)
}