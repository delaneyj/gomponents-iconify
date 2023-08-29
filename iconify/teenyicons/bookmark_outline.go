package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m12.5 14.5l-.312.39A.5.5 0 0 0 13 14.5h-.5Zm0-14h.5V0h-.5v.5Zm-10 0V0H2v.5h.5Zm0 14H2a.5.5 0 0 0 .812.39L2.5 14.5Zm5-4l.312-.39a.5.5 0 0 0-.624 0l.312.39Zm5.5 4V.5h-1v14h1ZM2 .5v14h1V.5H2Zm.812 14.39l5-4l-.624-.78l-5 4l.624.78Zm4.376-4l5 4l.624-.78l-5-4l-.624.78ZM12.5 0h-10v1h10V0Z"/>`),
		g.Group(children),
	)
}