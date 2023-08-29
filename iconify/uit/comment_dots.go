package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11.25a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5zm-3 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5zm6 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5zm4.415-5.96C15.71 1.195 9.385.88 5.29 4.584C1.195 8.289.88 14.614 4.584 18.709l-2.438 2.437A.5.5 0 0 0 2.5 22H12a10 10 0 0 0 6.709-2.585c4.096-3.705 4.412-10.03.706-14.125zM12 21H3.707l1.929-1.929a.5.5 0 0 0 0-.707a8.999 8.999 0 0 1 6.362-15.362A8.999 8.999 0 0 1 12 21z"/>`),
		g.Group(children),
	)
}