package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueReopened(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.28 5.656L2 7.006l-.66-.26L0 3.506l.92-.38l.81 1.95a6.48 6.48 0 0 1 12.48 1.93h-1a5.48 5.48 0 0 0-10.64-1.28l2.32-1l.39.93Zm8.86 2.68l1.34 3.23l-.92.44l-.82-2a6.49 6.49 0 0 1-12.5-2h1v-.5a5.49 5.49 0 0 0 10.64 1.89l-2.25.93l-.39-.92l3.25-1.35l.65.28Z" clip-rule="evenodd"/><circle cx="7.74" cy="7.54" r="1"/></g>`),
		g.Group(children),
	)
}