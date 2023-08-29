package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeutralFaceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2a16 16 0 1 0 16 16A16 16 0 0 0 18 2Zm7.05 21.06a1 1 0 0 1-1 1h-12a1 1 0 0 1 0-2h12a1 1 0 0 1 1 1ZM27 14.28a1.8 1.8 0 1 1-1.8-1.8a1.8 1.8 0 0 1 1.8 1.8Zm-15.8 1.8a1.8 1.8 0 1 1 1.8-1.8a1.8 1.8 0 0 1-1.84 1.8Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}