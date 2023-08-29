package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmTelehealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 25h-3v-3h-2v3h-3v2h3v3h2v-3h3v-2z"/><path fill="currentColor" d="M22 30h-2v-5c0-2.8-2.2-5-5-5H9c-2.8 0-5 2.2-5 5v5H2v-5c0-3.9 3.1-7 7-7h6c3.9 0 7 3.1 7 7v5zm6-24.8V3c0-1.1-.9-2-2-2h-8c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V6.8L31 9V3l-3 2.2zM26 9h-8V3h8v6zm-10.5 3.5c-.9.9-2.1 1.5-3.5 1.5c-2.8 0-5-2.2-5-5s2.2-5 5-5c.5 0 .9.1 1.3.2l.6-1.9c-.6-.2-1.2-.3-1.9-.3c-3.9 0-7 3.1-7 7s3.1 7 7 7c1.9 0 3.7-.8 5-2l-1.5-1.5z"/>`),
		g.Group(children),
	)
}