package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.98 5.004c-2.193 0-3.976 1.793-3.976 3.996c0 .732.211 1.409.555 2H8.354a3.942 3.942 0 0 0 .572-2.026c0-2.19-1.776-3.973-3.958-3.973c-2.182 0-3.957 1.782-3.957 3.973s1.775 3.973 3.957 3.973c.011 0 .031.053.031.053h8a4.001 4.001 0 0 0 3.958-4c0-2.203-1.784-3.996-3.977-3.996zM5 11c-1.104 0-2-.897-2-2s.896-2 2-2c1.103 0 2 .897 2 2s-.897 2-2 2zm8 0c-1.103 0-2-.897-2-1.999C11 7.898 11.898 7 13 7a2 2 0 0 1 0 4z"/>`),
		g.Group(children),
	)
}