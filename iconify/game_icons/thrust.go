package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thrust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M452.798 479.462L59.205 481.31s154.378-369.683 151.153-379.303l-46.12-.224l93.677-71.09l95.207 72.397l-50.263.224z"/>`),
		g.Group(children),
	)
}