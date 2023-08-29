package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.248 11.601c.45.313 1.05.313 1.5 0l9.088-5.281a.375.375 0 0 0-.048-.646l-9.205-3.537a1.315 1.315 0 0 0-1.17 0L.208 5.674a.375.375 0 0 0-.048.646l9.088 5.281zm10.54-.79l-2.486-1.233l-5.725 3.327c-.469.309-1.014.471-1.579.471s-1.11-.163-1.579-.471L2.698 9.576L.208 10.81a.375.375 0 0 0-.048.646l9.088 6.309c.45.313 1.05.313 1.5 0l9.088-6.309a.374.374 0 0 0-.048-.645z"/>`),
		g.Group(children),
	)
}