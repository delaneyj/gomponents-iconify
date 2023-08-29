package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudPakIntegration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.121 9.879a3 3 0 0 0-4.242 0l-.465.464l1.414 1.414l.465-.464a1 1 0 0 1 1.414 1.414L17.414 16l1.414 1.414l3.293-3.293a3 3 0 0 0 0-4.242zm-9.414 10.828a1 1 0 0 1-1.414-1.414L14.586 16l-1.414-1.414l-3.293 3.293a3 3 0 1 0 4.242 4.242l.465-.464l-1.414-1.414zM14.121 9.88a3 3 0 0 0-4.241 4.24l.464.465l1.414-1.414l-.464-.465a1 1 0 0 1 1.414-1.414L16 14.586l1.414-1.414z"/><path fill="currentColor" d="m22.121 17.879l-.464-.465l-1.414 1.414l.464.465a1 1 0 0 1-1.414 1.414L16 17.414l-1.414 1.414l3.293 3.293a3 3 0 0 0 4.242-4.242Z"/><path fill="currentColor" d="M16 31a.999.999 0 0 1-.504-.136l-12-7A1 1 0 0 1 3 23V9a1 1 0 0 1 .496-.864l12-7a1 1 0 0 1 1.008 0l12 7l-1.008 1.728L16 3.158L5 9.574v12.852l11 6.417l11-6.417V15h2v8a1 1 0 0 1-.496.864l-12 7A.999.999 0 0 1 16 31Z"/>`),
		g.Group(children),
	)
}