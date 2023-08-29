package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Untappd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.802 6.411c-1.585-.918-2.48-.387-1.89.337c-1.01.12-.408.891-.408.891s-5.272 6.779-7.293 9.245c-2.1 2.564-2.847 1.057-4.403 3.744L4.5 36.705s.168 2.332 3.516 4.27s5.116.728 5.116.728l9.308-16.077c1.556-2.687-.123-2.584 1.055-5.682c1.133-2.98 4.386-10.928 4.386-10.928s.969.139.57-.798c.922.152.937-.888-.649-1.807Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.593 12.638c-.822-2.028-1.475-3.622-1.475-3.622s-.968.139-.57-.798c-.922.152-.937-.889.65-1.807s2.48-.387 1.89.337c1.01.12.407.891.407.891L24 9.569m2.407 3.069a328.368 328.368 0 0 0 3.382 4.246c2.1 2.564 2.847 1.057 4.403 3.744L43.5 36.705s-.168 2.332-3.516 4.27s-5.116.728-5.116.728L25.56 25.626c-1.556-2.687.123-2.584-1.055-5.682c-.14-.367-.311-.81-.505-1.302"/>`),
		g.Group(children),
	)
}