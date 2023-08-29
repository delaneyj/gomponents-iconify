package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wpexplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 12A12 12 0 1 1 12 0a12.008 12.008 0 0 1 12 12Zm-1.5 0A10.5 10.5 0 1 0 12 22.5A10.516 10.516 0 0 0 22.5 12ZM7.542 5.841l4.074 1.739l-1.739 4.073L5.8 9.914l1.742-4.073Zm5.158 7.926l2.185 4.406H14.2l-2.343-4.687l-2.295 4.687h-.656l2.4-5.01l-1.046-.441l.282-.656l3.215 1.364l-.281.67Zm-.553-5.451l3.216 1.378l-1.378 3.2l-3.2-1.364l1.364-3.215Zm3.764 2.011l2.56 1.082l-1.1 2.546l-2.545-1.083l1.082-2.545Z"/>`),
		g.Group(children),
	)
}