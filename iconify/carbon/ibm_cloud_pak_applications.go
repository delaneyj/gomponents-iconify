package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudPakApplications(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 22h-6a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h6a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2Zm-6-6v4h6v-4Z"/><path fill="currentColor" d="M11 17H9v-5a2.002 2.002 0 0 1 2-2h6v2h-6Z"/><path fill="currentColor" d="M16 31a.999.999 0 0 1-.504-.136l-12-7A1 1 0 0 1 3 23V9a1 1 0 0 1 .496-.864l12-7a1 1 0 0 1 1.008 0l12 7l-1.008 1.728L16 3.158L5 9.574v12.852l11 6.417l11-6.417V15h2v8a1 1 0 0 1-.496.864l-12 7A.999.999 0 0 1 16 31Z"/>`),
		g.Group(children),
	)
}