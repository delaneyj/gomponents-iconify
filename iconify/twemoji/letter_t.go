package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterT(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M15.676 11.203h-3.38c-1.488 0-2.108-1.085-2.108-2.139c0-1.085.775-2.14 2.108-2.14h11.411c1.332 0 2.108 1.054 2.108 2.14c0 1.054-.619 2.139-2.108 2.139h-3.381v15.565c0 1.551-.992 2.418-2.325 2.418c-1.333 0-2.325-.867-2.325-2.418V11.203z"/>`),
		g.Group(children),
	)
}