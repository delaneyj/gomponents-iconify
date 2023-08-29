package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M9.925 9.032c0-1.271.93-2.294 2.325-2.294c1.333 0 2.325.868 2.325 2.294v6.697l7.627-8.124c.342-.372.93-.868 1.799-.868c1.178 0 2.295.899 2.295 2.232c0 .806-.496 1.457-1.52 2.48l-5.861 5.767l7.162 7.473c.744.744 1.303 1.426 1.303 2.357c0 1.457-1.146 2.139-2.418 2.139c-.898 0-1.488-.526-2.357-1.457l-8.031-8.682v7.906c0 1.21-.93 2.232-2.325 2.232c-1.333 0-2.325-.867-2.325-2.232V9.032z"/>`),
		g.Group(children),
	)
}