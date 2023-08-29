package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Management(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 128v288l96-96l96 96V128h128v768H320V128h256zm-448 0h128v768H128V128z"/>`),
		g.Group(children),
	)
}