package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkOutlineWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 24a10 10 0 1 1 10-10h2a12 12 0 1 0-12 12Z"/><path fill="currentColor" d="M12 15.59L9.41 13L8 14.41l4 4l7-7L17.59 10L12 15.59zM27.38 28h-6.762L24 21.236zM24 18a1 1 0 0 0-.895.553l-5 10A1 1 0 0 0 19 30h10a1 1 0 0 0 .921-1.39l-5.026-10.057A1 1 0 0 0 24 18z"/>`),
		g.Group(children),
	)
}