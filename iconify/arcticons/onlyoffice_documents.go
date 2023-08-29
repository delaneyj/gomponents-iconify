package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnlyofficeDocuments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.358 12.581L26.363 6.04a5.738 5.738 0 0 0-4.857-.001l-13.863 6.47a2.295 2.295 0 0 0-.001 4.159l13.995 6.541a5.738 5.738 0 0 0 4.857.002l13.863-6.471a2.295 2.295 0 0 0 .001-4.159Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.227 19.278l-5.584 2.606a2.295 2.295 0 0 0-.001 4.16l13.995 6.54a5.738 5.738 0 0 0 4.857.002l13.863-6.47a2.295 2.295 0 0 0 .001-4.159l-5.585-2.61"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.227 28.654L7.643 31.26a2.295 2.295 0 0 0-.001 4.16l13.995 6.54a5.738 5.738 0 0 0 4.857.001l13.863-6.47a2.295 2.295 0 0 0 .001-4.159l-5.585-2.61"/>`),
		g.Group(children),
	)
}