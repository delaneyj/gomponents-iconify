package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.224-.447a.5.5 0 0 0-.448 0L7.5.5Zm-6 3l-.224-.447A.5.5 0 0 0 1 3.5h.5Zm12 0h.5a.5.5 0 0 0-.276-.447L13.5 3.5Zm-8 7V10H5v.5h.5Zm4 0h.5V10h-.5v.5ZM0 15h15v-1H0v1ZM7.276.053l-6 3l.448.894l6-3l-.448-.894Zm6.448 3l-6-3l-.448.894l6 3l.448-.894ZM7 3v2.5h1V3H7Zm0 2.5V8h1V5.5H7ZM5 6h2.5V5H5v1Zm2.5 0H10V5H7.5v1ZM1 3.5v11h1v-11H1Zm12 0v11h1v-11h-1Zm-7 11v-4H5v4h1ZM5.5 11h4v-1h-4v1Zm3.5-.5v4h1v-4H9Z"/>`),
		g.Group(children),
	)
}