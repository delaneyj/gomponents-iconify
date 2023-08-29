package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkOutlineError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 24a10 10 0 1 1 10-10h2a12 12 0 1 0-12 12Z"/><path fill="currentColor" d="M12 15.59L9.41 13L8 14.41l4 4l7-7L17.59 10L12 15.59zM30 24a6 6 0 1 0-6 6a6.007 6.007 0 0 0 6-6zm-2 0a3.952 3.952 0 0 1-.567 2.019l-5.452-5.452A3.953 3.953 0 0 1 24 20a4.005 4.005 0 0 1 4 4zm-8 0a3.952 3.952 0 0 1 .567-2.019l5.452 5.452A3.953 3.953 0 0 1 24 28a4.005 4.005 0 0 1-4-4z"/>`),
		g.Group(children),
	)
}