package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ambulance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.18 14a2 2 0 1 1-3.999.001A2 2 0 0 1 6.18 14zM14 14a2 2 0 1 1-3.999.001A2 2 0 0 1 14 14zM5 6H4v1H3v1h1v1h1V8h1V7H5V6z"/><path fill="currentColor" d="m15.76 8.64l-3-4.53A2.501 2.501 0 0 0 10.682 3H8V2a1 1 0 0 0-2 0v1H1.5A1.5 1.5 0 0 0 0 4.5V13h1.37a3.238 3.238 0 0 1 2.812-2a3.236 3.236 0 0 1 2.81 1.978l2.188.021a3.238 3.238 0 0 1 2.812-2a3.003 3.003 0 0 1 2.822 1.979l1.187.021v-3.57a1.427 1.427 0 0 0-.243-.795zm-8.84-.52a2.5 2.5 0 1 1-3.017-2.997a2.48 2.48 0 0 1 3.014 3.014zM10 8V5h.859a2.25 2.25 0 0 1 1.866.992L14.05 8z"/>`),
		g.Group(children),
	)
}