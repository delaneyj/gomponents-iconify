package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePocket0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePocket1" fill="currentColor" fill-rule="nonzero"><path id="fePocket2" d="M21.902 4.194A1.82 1.82 0 0 0 20.197 3H3.813A1.823 1.823 0 0 0 2 4.814v6.035l.069 1.2c.29 2.73 1.707 5.116 3.9 6.779c.038.03.078.059.118.089l.025.018a9.897 9.897 0 0 0 3.91 1.727a10.06 10.06 0 0 0 4.05-.014a.261.261 0 0 0 .064-.023a9.906 9.906 0 0 0 3.753-1.691l.025-.018c.04-.03.08-.058.119-.09c2.192-1.663 3.609-4.048 3.898-6.778l.069-1.2V4.814a1.792 1.792 0 0 0-.098-.62Zm-4.235 6.287l-4.704 4.513a1.372 1.372 0 0 1-1.899 0L6.36 10.48a1.371 1.371 0 1 1 1.898-1.979l3.756 3.601l3.756-3.6a1.372 1.372 0 0 1 1.898 1.978Z"/></g></g>`),
		g.Group(children),
	)
}