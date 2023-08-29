package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modsawup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M36 208c-2.21 0-2.619-1.144-.926-2.546L220.926 51.546c1.698-1.406 3.074-.76 3.074 1.464v150.98a4.003 4.003 0 0 1-3.996 4.01h-8.008a3.996 3.996 0 0 1-3.996-4.007V84.007c0-2.213-1.387-2.861-3.079-1.465L56.08 205.458C54.379 206.862 51.209 208 49 208H35.999z"/>`),
		g.Group(children),
	)
}