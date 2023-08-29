package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.5 9h-1v7h1c.275 0 .5-.225.5-.5v-6c0-.275-.225-.5-.5-.5zm7 0c-.275 0-.5.225-.5.5v6c0 .275.225.5.5.5h1V9h-1z"/><path fill="currentColor" d="M16 8A8 8 0 1 0 .479 10.732A3.5 3.5 0 0 0 3 15.964V9.036a3.478 3.478 0 0 0-1.371.506a6.5 6.5 0 1 1 12.743 0A3.484 3.484 0 0 0 13 9.035v6.929a3.5 3.5 0 0 0 2.521-5.232C15.831 9.879 16 8.959 16 8z"/>`),
		g.Group(children),
	)
}