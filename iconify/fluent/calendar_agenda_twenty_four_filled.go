package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarAgendaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 17.75A3.25 3.25 0 0 1 17.75 21H6.25A3.25 3.25 0 0 1 3 17.75V6.25A3.25 3.25 0 0 1 6.25 3h11.5A3.25 3.25 0 0 1 21 6.25v11.5Zm-4-10a.75.75 0 0 0-.648-.743L16.25 7h-8.5l-.102.007a.75.75 0 0 0 0 1.486l.102.007h8.5l.102-.007A.75.75 0 0 0 17 7.75Zm0 8.5a.75.75 0 0 0-.648-.743l-.102-.007h-8.5l-.102.007a.75.75 0 0 0 0 1.486L7.75 17h8.5l.102-.007A.75.75 0 0 0 17 16.25ZM17 12a.75.75 0 0 0-.648-.743l-.102-.007h-8.5l-.102.007a.75.75 0 0 0 0 1.486l.102.007h8.5l.102-.007A.75.75 0 0 0 17 12Z"/>`),
		g.Group(children),
	)
}