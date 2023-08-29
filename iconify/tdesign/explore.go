package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Explore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.364 5.636A9 9 0 1 0 5.636 18.364A9 9 0 0 0 18.364 5.636ZM4.222 4.222c4.296-4.296 11.26-4.296 15.556 0c4.296 4.296 4.296 11.26 0 15.556c-4.296 4.296-11.26 4.296-15.556 0c-4.296-4.296-4.296-11.26 0-15.556Zm13.22 2.337l-4.965 12.91l-2.1-5.844l-5.845-2.1l12.91-4.966Zm-7.174 4.902l1.672.6l.6 1.672l1.42-3.692l-3.692 1.42Z"/>`),
		g.Group(children),
	)
}