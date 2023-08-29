package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSunOne0"><g fill="#fff"><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 37c7.18 0 13-5.82 13-13s-5.82-13-13-13s-13 5.82-13 13s5.82 13 13 13Z"/><path d="M24 6a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm14.5 6a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm6 14.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-6 14.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM24 47a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM9.5 41a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm-6-14.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm6-14.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSunOne0)"/>`),
		g.Group(children),
	)
}