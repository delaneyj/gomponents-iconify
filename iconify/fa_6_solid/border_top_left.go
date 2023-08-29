package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 448c0 17.7 14.3 32 32 32s32-14.3 32-32V112c0-8.8 7.2-16 16-16h336c17.7 0 32-14.3 32-32s-14.3-32-32-32H80C35.8 32 0 67.8 0 112v336zm160 0a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm192 0a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm-96 0a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm192 0a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm-32-160a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm0 32a32 32 0 1 0 0 64a32 32 0 1 0 0-64zm0-128a32 32 0 1 0 0-64a32 32 0 1 0 0 64z"/>`),
		g.Group(children),
	)
}