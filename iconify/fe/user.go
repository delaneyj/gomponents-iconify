package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func User(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feUser0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feUser1" fill="currentColor" fill-rule="nonzero"><path id="feUser2" d="M12 12a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0 3c3.186 0 6.045.571 8 3.063V20H4v-1.937C5.955 15.57 8.814 15 12 15Z"/></g></g>`),
		g.Group(children),
	)
}