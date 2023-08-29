package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Git(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feGit0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feGit1" fill="currentColor"><path id="feGit2" d="M21.623 11.11L12.89 2.376a1.288 1.288 0 0 0-1.821 0L9.256 4.191l2.3 2.3a1.53 1.53 0 0 1 1.937 1.95l2.217 2.217a1.532 1.532 0 1 1-.918.864l-2.068-2.068v5.441a1.533 1.533 0 1 1-1.26-.045V9.36a1.532 1.532 0 0 1-.832-2.01L8.365 5.081l-5.988 5.987a1.289 1.289 0 0 0 0 1.822l8.733 8.732a1.288 1.288 0 0 0 1.821 0l8.692-8.692a1.288 1.288 0 0 0 0-1.822"/></g></g>`),
		g.Group(children),
	)
}