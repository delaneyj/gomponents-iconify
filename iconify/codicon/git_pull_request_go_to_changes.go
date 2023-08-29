package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequestGoToChanges(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3 10v4l1 1h9l1-1V5l-.29-.71l-3-3L10 1H8v1h2l3 3v9H4v-4H3Zm8-4H9V4H8v2H6v1h2v2h1V7h2V6Zm-5 5h5v1H6v-1Z"/><path d="M7.06 3.854L4.915 6l-.707-.707L5.5 4h-3a1.5 1.5 0 0 0 0 3H3v1h-.5a2.5 2.5 0 1 1 0-5h3L4.207 1.707L4.914 1l2.147 2.146v.708Z"/></g>`),
		g.Group(children),
	)
}