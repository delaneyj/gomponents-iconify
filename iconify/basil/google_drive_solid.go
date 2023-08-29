package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.243 21.32a.3.3 0 0 1-.26-.45l3.194-5.507a.3.3 0 0 1 .259-.149l11.62-.001a.3.3 0 0 1 .26.45l-3.178 5.507a.3.3 0 0 1-.26.15H7.243Zm.094-16.348a.3.3 0 0 1 .52 0l3.185 5.493a.3.3 0 0 1 0 .3L5.21 20.87a.3.3 0 0 1-.52 0l-3.178-5.506a.3.3 0 0 1 0-.3L7.337 4.972ZM21.47 13.3a.3.3 0 0 1-.26.45h-6.355a.3.3 0 0 1-.26-.15L8.736 3.45a.3.3 0 0 1 .26-.45h6.355a.3.3 0 0 1 .26.15l5.86 10.15Z"/>`),
		g.Group(children),
	)
}