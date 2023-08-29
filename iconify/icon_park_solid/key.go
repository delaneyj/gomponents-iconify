package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSKey0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M22.868 24.298a9.865 9.865 0 0 1 2.63 9.588a9.926 9.926 0 0 1-7.066 7.028a9.988 9.988 0 0 1-9.638-2.615a9.863 9.863 0 0 1 .121-13.878c3.839-3.82 10.046-3.873 13.951-.121l.002-.002Z"/><path stroke-linecap="round" d="M23 24L40 7"/><path fill="#fff" d="m30.305 16.9l5.429 5.4l6.333-6.3l-5.428-5.4l-6.334 6.3Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSKey0)"/>`),
		g.Group(children),
	)
}