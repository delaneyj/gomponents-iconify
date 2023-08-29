package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Platformdotsh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 0H0v9.541h24V0zm0 20.755H0V24h24v-3.245zM0 12.618h24v4.892H0v-4.892z"/>`),
		g.Group(children),
	)
}