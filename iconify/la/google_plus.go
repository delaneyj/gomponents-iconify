package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 7c-4.973 0-9 4.027-9 9s4.027 9 9 9s9-4.027 9-9c0-.617-.066-1.219-.188-1.8l-.046-.2H11v3h6c-.477 2.836-3.027 5-6 5c-3.313 0-6-2.688-6-6c0-3.313 2.688-6 6-6c1.5 0 2.867.555 3.922 1.465l2.148-2.106A8.963 8.963 0 0 0 11 7zm14 4v3h-3v2h3v3h2v-3h3v-2h-3v-3z"/>`),
		g.Group(children),
	)
}