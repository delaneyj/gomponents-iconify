package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadFaceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm9 12.28a1.8 1.8 0 1 1-1.8-1.8a1.8 1.8 0 0 1 1.8 1.8Zm-15.55 1.8a1.8 1.8 0 1 1 1.8-1.8a1.8 1.8 0 0 1-1.84 1.8Zm14 7.53a1 1 0 0 1-1.6 1.2a7 7 0 0 0-11.31.13a1 1 0 1 1-1.63-1.16a9 9 0 0 1 14.54-.17Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}