package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentUnknownTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M10 2v4.5A1.5 1.5 0 0 0 11.5 8H16v8.5a1.5 1.5 0 0 1-1.5 1.5H9.743A5.5 5.5 0 0 0 4 9.207V3.5A1.5 1.5 0 0 1 5.5 2H10z" fill="currentColor"/><path d="M11 2.25V6.5a.5.5 0 0 0 .5.5h4.25L11 2.25z" fill="currentColor"/><path d="M10 14.5a4.5 4.5 0 1 0-9 0a4.5 4.5 0 0 0 9 0zm-4.5 1.88a.625.625 0 1 1 0 1.25a.625.625 0 0 1 0-1.25zm0-4.877c1.031 0 1.853.846 1.853 1.95c0 .586-.214.908-.727 1.319l-.277.214c-.246.194-.329.3-.346.448l-.011.156A.5.5 0 0 1 5 15.5c0-.57.21-.884.716-1.288l.278-.215c.288-.23.36-.342.36-.544c0-.558-.382-.95-.854-.95c-.494 0-.859.366-.853.945a.5.5 0 1 1-1 .01c-.011-1.137.805-1.955 1.853-1.955z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}