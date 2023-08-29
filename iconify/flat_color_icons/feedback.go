package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feedback(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#78909C" d="M40 41H8c-2.2 0-4-1.8-4-4V16.1c0-1.3.6-2.5 1.7-3.3L24 0l18.3 12.8c1.1.7 1.7 2 1.7 3.3V37c0 2.2-1.8 4-4 4z"/><path fill="#fff" d="M12 11h24v22H12z"/><path fill="#9C27B0" d="m24 13.6l-6 7.8h12z"/><path fill="#CFD8DC" d="M40 41H8c-2.2 0-4-1.8-4-4V17l20 13l20-13v20c0 2.2-1.8 4-4 4z"/><path fill="#9C27B0" d="m24 28l2-1.3V20h-4v6.7z"/>`),
		g.Group(children),
	)
}