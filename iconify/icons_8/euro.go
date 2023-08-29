package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 4c-4.738 0-8.587 3.887-9.688 9H6v2h2.063c-.023.328-.063.666-.063 1c0 .334.04.672.063 1H6v2h2.313c1.1 5.113 4.95 9 9.687 9c2.71 0 5.17-1.303 6.938-3.344l-1.532-1.312C21.954 25.02 20.07 26 18 26c-3.502 0-6.59-2.898-7.625-7H19v-2h-8.97c-.025-.33-.03-.66-.03-1c0-.34.005-.67.03-1H19v-2h-8.625C11.41 8.898 14.498 6 18 6c2.07 0 3.954.98 5.406 2.656l1.532-1.312C23.168 5.304 20.708 4 18 4z"/>`),
		g.Group(children),
	)
}