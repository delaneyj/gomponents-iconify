package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligiousShinto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.5 2.5h-10a.945.945 0 0 0-1 1a.945.945 0 0 0 1 1H3V12a.945.945 0 0 0 1 1a.945.945 0 0 0 1-1V8h5v4a.945.945 0 0 0 1 1a.945.945 0 0 0 1-1V4.5h.5a.945.945 0 0 0 1-1a.945.945 0 0 0-1-1ZM10 6H8V4.5h2ZM7 6H5V4.5h2Z"/>`),
		g.Group(children),
	)
}