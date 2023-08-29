package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequestNewChanges(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m13.71 4.29l-3-3L10 1H4L3 2v12l1 1h5.354a4.019 4.019 0 0 1-.819-1H4V2h6l3 3v3.126c.355.091.69.23 1 .41V5l-.29-.71ZM8.126 11H6v1h2c0-.345.044-.68.126-1ZM6 6h2V4h1v2h2v1H9v2H8V7H6V6Z" clip-rule="evenodd"/><path d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/></g>`),
		g.Group(children),
	)
}