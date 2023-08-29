package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M32 128h960v64H32z"/><path fill="currentColor" d="M192 192v512h640V192H192zm-64-64h768v608a32 32 0 0 1-32 32H160a32 32 0 0 1-32-32V128z"/><path fill="currentColor" d="M322.176 960H248.32l144.64-250.56l55.424 32L322.176 960zm453.888 0h-73.856L576 741.44l55.424-32L776.064 960z"/>`),
		g.Group(children),
	)
}