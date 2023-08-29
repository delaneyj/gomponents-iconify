package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseEtcd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.944 20.942v-6.6h-1.77v-2.345h.92c.874 0 1.15-.414 1.15-1.242V8.708h2.644v3.289h2.46v2.345h-2.46v7.313h2.277V24h-2.116a2.778 2.778 0 0 1-3.105-3.058zM6.684 24V7.95h10.579v2.69H9.72v3.886h6.669v2.69H9.72v4.094h7.543V24z"/>`),
		g.Group(children),
	)
}