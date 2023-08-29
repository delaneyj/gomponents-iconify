package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Platform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 832v-64h128v64h192v64H256v-64h192zM128 704V128h768v576H128z"/>`),
		g.Group(children),
	)
}