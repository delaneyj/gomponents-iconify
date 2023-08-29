package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailHasMail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M3.004 0v7l5.027 4l4.924-4V3.031L8 3.026V0H3.004zm7.965 8.016h-6V6.969h6v1.047zM8.953 5v1h-4V5h4z"/><path d="M9 0v1.951h3.848L9 0zm6.34 4.963l-2.566-1.896v1.105l2.32 1.986v.654l-3.031 2.789L15.079 14l-1.094.015l-2.59-3.739l-3.394 2.786l-3.262-2.772l-2.707 3.726H.923l3.135-4.414l-3.15-2.82l.015-.69L5.19 2.251L5.174.892L.725 4.964c-.359 0-.707.291-.707.65v8.649c0 .36.348.706.707.706h14.617c.359 0 .65-.346.65-.706V5.614a.654.654 0 0 0-.652-.651z"/></g>`),
		g.Group(children),
	)
}