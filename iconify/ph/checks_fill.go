package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChecksFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 48H32a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h192a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16ZM80 160a8 8 0 0 1-5.66-2.34l-24-24a8 8 0 0 1 11.32-11.32L80 140.69l50.34-50.35a8 8 0 0 1 11.32 11.32l-56 56A8 8 0 0 1 80 160Zm133.66-58.34l-56 56a8 8 0 0 1-11.32 0l-16-16a8 8 0 0 1 11.32-11.32L152 140.69l50.34-50.35a8 8 0 0 1 11.32 11.32Z"/>`),
		g.Group(children),
	)
}