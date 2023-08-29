package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.986 5.424a53.892 53.892 0 0 0-11.972 0a.343.343 0 0 0-.228.556l3.517 4.348a8.75 8.75 0 0 1 1.947 5.503v2.889l1.5 1.1v-3.99a8.75 8.75 0 0 1 1.947-5.502l3.517-4.348a.343.343 0 0 0-.228-.556ZM5.848 3.933a55.39 55.39 0 0 1 12.305 0c1.446.162 2.143 1.858 1.228 2.99l-3.518 4.348a7.25 7.25 0 0 0-1.613 4.56V21.3a.75.75 0 0 1-1.194.605l-3-2.2a.75.75 0 0 1-.306-.605v-3.27c0-1.66-.57-3.269-1.613-4.56L4.62 6.924c-.916-1.132-.22-2.828 1.228-2.99Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}