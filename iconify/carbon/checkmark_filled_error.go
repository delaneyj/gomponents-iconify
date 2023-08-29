package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkFilledError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 24a6 6 0 1 0-6 6a6.007 6.007 0 0 0 6-6Zm-2 0a3.952 3.952 0 0 1-.567 2.019l-5.452-5.452A3.953 3.953 0 0 1 24 20a4.005 4.005 0 0 1 4 4Zm-8 0a3.952 3.952 0 0 1 .567-2.019l5.452 5.452A3.953 3.953 0 0 1 24 28a4.005 4.005 0 0 1-4-4Z"/><path fill="currentColor" d="M14 2a12 12 0 1 0 2 23.82V24a8 8 0 0 1 8-8h1.82A11.935 11.935 0 0 0 14 2Zm-2 16.59l-4-4L9.59 13L12 15.41L17.41 10L19 11.59Z"/><path fill="none" d="m12 18.591l-4-4L9.591 13L12 15.409L17.409 10L19 11.591l-7 7z"/>`),
		g.Group(children),
	)
}