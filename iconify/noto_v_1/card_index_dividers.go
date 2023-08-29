package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardIndexDividers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#a13300" d="M120.7 17.6H85.3c-2.4 0-4.3 1.9-4.3 4.3v13.4h44V21.9c0-2.4-1.9-4.3-4.3-4.3z"/><path fill="#ccb656" d="M121.1 102.3H19.9c-2.1 0-3.9-1.7-3.9-3.9V31.7c0-2.1 1.7-3.9 3.9-3.9H125v70.7c0 2.1-1.7 3.8-3.9 3.8z"/><path fill="#1a6ca2" d="M83.2 25.3H47.7c-2.4 0-4.3 1.9-4.3 4.3V43h44V29.6c.1-2.4-1.9-4.3-4.2-4.3z"/><path fill="#e5d07a" d="M115.1 111H13.9c-2.1 0-3.9-1.7-3.9-3.9V40.3c0-2.1 1.7-3.9 3.9-3.9h101.2c2.1 0 3.9 1.7 3.9 3.9v66.8c0 2.1-1.7 3.9-3.9 3.9z"/><path fill="#f79329" d="M42.8 34.3H7.2c-2.3 0-4.2 1.9-4.2 4.2v13.2h44V38.5c0-2.3-1.9-4.2-4.2-4.2z"/><path fill="#f7e599" d="M108.1 120H6.9c-2.1 0-3.9-1.7-3.9-3.9V45.5h105.1c2.1 0 3.9 1.7 3.9 3.9v66.8c0 2.1-1.7 3.8-3.9 3.8z"/>`),
		g.Group(children),
	)
}