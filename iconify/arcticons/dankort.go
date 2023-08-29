package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dankort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.976 12.876a11.334 11.334 0 0 0-7.825 3.097h13.39c4.781 0 7.01 2.213 7.623 5.85c1.53-1.77 3.097-3.582 5.35-5.865l8.75-.126a11.052 11.052 0 0 0-8.047-2.956Zm28.353 5.02l-5.563 5.334l5.953 6.54a8.741 8.741 0 0 0 1.781-5.3a10.986 10.986 0 0 0-2.17-6.575ZM26.251 25.01a7.73 7.73 0 0 1-7.709 6.347H4.5a11.32 11.32 0 0 0 8.476 3.767h19.242a11.14 11.14 0 0 0 8.184-3.23l-8.829.033ZM9.997 21.96l-1.173 3.77l6.569-.098a2.708 2.708 0 0 0 2.106-1.853a1.508 1.508 0 0 0-1.286-1.8Z"/>`),
		g.Group(children),
	)
}