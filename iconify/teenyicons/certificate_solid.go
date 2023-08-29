package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5v4l-.8-.601a4.5 4.5 0 0 0-6.3 6.3l.1.134V15H1.5A1.5 1.5 0 0 1 0 13.5v-12ZM8 5H3V4h5v1ZM3 8h3V7H3v1Z"/><path d="M11.5 5A3.5 3.5 0 0 0 9 10.95v3.55a.5.5 0 0 0 .8.4l1.7-1.275l1.7 1.275a.5.5 0 0 0 .8-.4v-3.55A3.5 3.5 0 0 0 11.5 5ZM10 13.5v-1.837c.455.216.963.337 1.5.337s1.045-.12 1.5-.337V13.5l-1.2-.9a.5.5 0 0 0-.6 0l-1.2.9Z"/></g>`),
		g.Group(children),
	)
}