package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfTriangleGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="88.472" cy="11.528" r="5.824" fill="currentColor"/><circle cx="50" cy="50" r="5.824" fill="currentColor"/><circle cx="88.472" cy="88.472" r="5.824" fill="currentColor"/><circle cx="11.528" cy="88.472" r="5.824" fill="currentColor"/><circle cx="11.528" cy="11.528" r="5.824" fill="currentColor"/><path fill="currentColor" d="M52 88h27v.236c0-1.986.421-3.598 1.438-5.146L55.446 58.17A9.773 9.773 0 0 1 52 59.618V88zm-32.394-4.909C20.623 84.638 21 86.249 21 88.236V88h27V59.618a9.746 9.746 0 0 1-3.351-1.391L19.606 83.091zm-2.697-63.485C15.362 20.623 13.751 21 11.764 21H12v27h28.382a9.746 9.746 0 0 1 1.391-3.351L16.909 19.606zM12 52v27h-.236c1.986 0 3.598.421 5.145 1.438l24.92-24.993a9.773 9.773 0 0 1-1.448-3.446H12zm68.394-35.09C79.377 15.362 79 13.751 79 11.764V12H52v28.382a9.746 9.746 0 0 1 3.351 1.391L80.394 16.91zM48 12H21v-.236c0 1.986-.421 3.598-1.438 5.145l24.993 24.92a9.773 9.773 0 0 1 3.446-1.448V12zm40 36V21h.236c-1.986 0-3.598-.421-5.145-1.438L58.17 44.554A9.773 9.773 0 0 1 59.618 48H88zm0 4H59.618a9.746 9.746 0 0 1-1.391 3.351L83.09 80.394C84.638 79.377 86.249 79 88.236 79H88V52z"/>`),
		g.Group(children),
	)
}