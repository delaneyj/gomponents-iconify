package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kcs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#0093DD"/><path fill="#FFF" d="m13.54 16l5.174 5.33l3.265-3.363a1.446 1.446 0 0 1 2.088 0a1.554 1.554 0 0 1 0 2.152l-4.309 4.44a1.456 1.456 0 0 1-2.088 0l-6.216-6.406v3.808c0 .836-.666 1.522-1.477 1.522c-.816 0-1.477-.682-1.477-1.522V10.04c0-.84.661-1.522 1.477-1.522c.815 0 1.477.682 1.477 1.522v3.808l6.216-6.406a1.456 1.456 0 0 1 2.088 0l4.31 4.44a1.554 1.554 0 0 1 0 2.152a1.446 1.446 0 0 1-2.09 0l-3.264-3.364L13.54 16zm5.176-1.523c.816 0 1.478.682 1.478 1.523c0 .841-.662 1.523-1.478 1.523s-1.478-.682-1.478-1.523c0-.841.662-1.523 1.478-1.523z"/></g>`),
		g.Group(children),
	)
}