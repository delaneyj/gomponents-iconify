package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 4h-4V3c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v1H1c-.55 0-1 .45-1 1v9c0 .55.45 1 1 1h14c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1zm-9-.998L6.002 3h3.996l.002.002V4H6v-.998zM15 8h-2v1.5c0 .275-.225.5-.5.5h-1a.501.501 0 0 1-.5-.5V8H5v1.5c0 .275-.225.5-.5.5h-1a.501.501 0 0 1-.5-.5V8H1V7h14v1z"/>`),
		g.Group(children),
	)
}