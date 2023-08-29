package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeniedSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm2.564-4.23a6.5 6.5 0 0 0 9.165 9.165L2.564 3.272Zm.707-.706l9.165 9.165a6.5 6.5 0 0 0-9.165-9.165Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}