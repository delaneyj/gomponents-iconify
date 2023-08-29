package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandCream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#2F88FF" d="M19 17C16.2688 18.1375 16.5121 21.5781 16.9998 23H30.9998C34.9016 17.3126 26.7838 11.1512 26.7838 13.521C26.7838 15.8907 22.414 15.5781 19 17Z"/><rect width="22" height="8" x="13" y="23"/></g>`),
		g.Group(children),
	)
}