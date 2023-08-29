package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopLetterCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M20 6.5H6a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1Zm-13 11v-9h12v9H7Z"/><path d="m20.648 8.261l-7.045 6a1 1 0 0 1-1.301-.004l-6.955-6C4.645 7.652 5.073 6.5 6 6.5h14c.93 0 1.356 1.158.648 1.761ZM8.69 8.5l4.27 3.683L17.282 8.5H8.69Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopLetterCircleFilled0)"/></g>`),
		g.Group(children),
	)
}