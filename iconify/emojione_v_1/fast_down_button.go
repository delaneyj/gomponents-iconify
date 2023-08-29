package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastDownButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M64 57.1a6.9 6.9 0 0 1-6.898 6.904H6.892A6.9 6.9 0 0 1-.004 57.1V6.9c0-3.81 3.088-6.9 6.896-6.9h50.21A6.899 6.899 0 0 1 64 6.9v50.2"/><g fill="#fff"><path d="M15.685 11.924c.156-2.318 1.171-4.253 2.557-5.101h27.674c1.39.852 2.404 2.791 2.559 5.113L32.134 34.369L15.685 11.924"/><path d="M15.685 34.901c.156-2.318 1.171-4.253 2.557-5.101h27.674c1.39.851 2.404 2.791 2.559 5.112L32.134 57.347L15.685 34.901"/></g>`),
		g.Group(children),
	)
}