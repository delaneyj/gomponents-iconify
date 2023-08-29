package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M64 57.1a6.899 6.899 0 0 1-6.896 6.903H6.894A6.9 6.9 0 0 1-.002 57.1V6.9A6.899 6.899 0 0 1 6.894.001h50.21C60.914.001 64 3.091 64 6.9v50.2"/><g fill="#fff"><path d="M11.899 15.1c-2.313.157-4.245 1.168-5.091 2.553v27.618c.85 1.388 2.786 2.4 5.103 2.554l22.39-16.303L11.899 15.1"/><path d="M34.831 15.1c-2.313.157-4.245 1.168-5.09 2.553v27.618c.85 1.388 2.785 2.4 5.103 2.554l22.39-16.303L34.831 15.1"/></g>`),
		g.Group(children),
	)
}