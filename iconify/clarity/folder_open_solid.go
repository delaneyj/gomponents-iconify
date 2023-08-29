package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M35.32 13.74a1.71 1.71 0 0 0-1.45-.74h-22.7a2.59 2.59 0 0 0-2.25 1.52a1 1 0 0 0 0 .14L6 25V7h6.49l2.61 3.59a1 1 0 0 0 .81.41H32a2 2 0 0 0-2-2H16.42l-2.31-3.18A2 2 0 0 0 12.49 5H6a2 2 0 0 0-2 2v22.69A1.37 1.37 0 0 0 5.41 31h24.93a1 1 0 0 0 1-.72l4.19-15.1a1.64 1.64 0 0 0-.21-1.44Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}