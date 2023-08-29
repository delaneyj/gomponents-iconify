package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.32 19.98c-.58 0-1.16-.22-1.6-.66l-6.48-6.47a.75.75 0 0 1-.22-.53V4.77c0-.41.34-.75.75-.75h7.54c.2 0 .39.08.53.22l6.48 6.48c.87.88.87 2.31 0 3.19l-5.41 5.41c-.44.44-1.02.66-1.6.66Zm-6.8-7.97l6.26 6.25c.3.29.78.29 1.07 0l5.41-5.41c.29-.29.29-.77 0-1.07L12 5.52H5.52V12Zm13.27 1.37ZM8.5 9.75a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/>`),
		g.Group(children),
	)
}