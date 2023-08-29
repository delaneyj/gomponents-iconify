package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 20a12 12 0 0 0-12 12v28H76V32a12 12 0 0 0-24 0v192a12 12 0 0 0 24 0v-28h104v28a12 12 0 0 0 24 0V32a12 12 0 0 0-12-12Zm-12 64v32H76V84ZM76 172v-32h104v32Z"/>`),
		g.Group(children),
	)
}