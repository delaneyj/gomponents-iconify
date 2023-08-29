package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudPakWatsonAiops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 24a1 1 0 0 1-.961-.725L11.246 17H8v-2h4a1 1 0 0 1 .961.725L14 19.36l3.039-10.635a1 1 0 0 1 1.922 0L20.754 15H24v2h-4a1 1 0 0 1-.961-.725L18 12.64l-3.039 10.635A1 1 0 0 1 14 24Z"/><path fill="currentColor" d="M16 31a.999.999 0 0 1-.504-.136l-12-7A1 1 0 0 1 3 23V9a1 1 0 0 1 .496-.864l12-7a1 1 0 0 1 1.008 0l12 7l-1.008 1.728L16 3.158L5 9.574v12.852l11 6.417l11-6.417V15h2v8a1 1 0 0 1-.496.864l-12 7A.999.999 0 0 1 16 31Z"/>`),
		g.Group(children),
	)
}