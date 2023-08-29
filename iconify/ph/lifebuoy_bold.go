package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LifebuoyBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm44.25 135.28a51.89 51.89 0 0 0 0-54.56l23-23a83.84 83.84 0 0 1 0 100.56ZM100 128a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm78.28-67.25l-23 23a51.89 51.89 0 0 0-54.56 0l-23-23a83.84 83.84 0 0 1 100.56 0Zm-117.53 17l23 23a51.89 51.89 0 0 0 0 54.56l-23 23a83.84 83.84 0 0 1 0-100.56Zm17 117.53l23-23a51.89 51.89 0 0 0 54.56 0l23 23a83.84 83.84 0 0 1-100.56 0Z"/>`),
		g.Group(children),
	)
}