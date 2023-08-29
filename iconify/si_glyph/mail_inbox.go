package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailInbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M6.031 4.017L9 6.97l2.969-2.953H9.984V1.032H8.047v2.985H6.031z"/><path d="M16.062 7.094v-1.37l-7.048 5.62l-7.076-5.532l.029 1.282l3.07 2.599l-3.189 4.353h1.246l2.862-3.477l3.058 2.619l3.008-2.545l2.884 3.403l1.264-.015l-3.124-4.338l3.016-2.599z"/><path fill="currentColor" d="M16.304 5.059h-3.242l-4 4.004l-4.103-4.004H1.701a.65.65 0 0 0-.648.648v8.617a.65.65 0 0 0 .648.648h14.603a.65.65 0 0 0 .649-.648V5.707a.65.65 0 0 0-.649-.648zm-1.398 8.987l-2.884-3.403l-3.009 2.545l-3.058-2.618l-2.862 3.477H1.847l3.189-4.353l-3.07-2.6l-.029-1.281l7.076 5.531l7.049-5.62v1.37l-3.017 2.6l3.124 4.338l-1.263.014z"/></g>`),
		g.Group(children),
	)
}