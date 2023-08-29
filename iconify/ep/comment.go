package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 504a56 56 0 1 1 0-112a56 56 0 0 1 0 112zm-224 0a56 56 0 1 1 0-112a56 56 0 0 1 0 112zm-224 0a56 56 0 1 1 0-112a56 56 0 0 1 0 112zM128 128v640h192v160l224-160h352V128H128z"/>`),
		g.Group(children),
	)
}