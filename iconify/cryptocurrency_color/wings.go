package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#0dc9f7"/><g fill="#fff" fill-rule="nonzero"><path fill-opacity=".305" d="m18.904 15.739l-3.045 2.364l-1.247-4.325l-7.224-1.935l9.481.487z"/><path fill-opacity=".7" d="M8.27 23.993L24.586 11.33L26 14.476l-1.855-.513l-.065 3.264z"/><path d="m22.796 17.78l-4.747-8.161L6 9l9.183 2.461l2.49 8.49z"/></g></g>`),
		g.Group(children),
	)
}