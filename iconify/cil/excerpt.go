package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Excerpt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 344h480v32H16zm0-191.667h480v32H16zM16 248h328v32H16zm0-192h384v32H16zm0 384h32v32H16zm224 0h32v32h-32zm-112 0h32v32h-32z"/>`),
		g.Group(children),
	)
}