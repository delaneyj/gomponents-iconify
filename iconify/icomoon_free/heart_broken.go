package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.8 1C14.118 1 16 2.882 16 5.2c0 4.566-4.935 5.982-8 10.616c-3.243-4.663-8-5.9-8-10.616C0 2.881 1.882 1 4.2 1c.943 0 1.812.43 2.512 1.06L5.499 4l3.5 2l-2 5l5.5-6l-3.5-2l.967-1.451c.553-.34 1.175-.549 1.833-.549z"/>`),
		g.Group(children),
	)
}