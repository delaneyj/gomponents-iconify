package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPathConnectLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M16.5 7.5h-3"/><path d="M5 5.217c0-.573 0-.86.049-1.099c.213-1.052 1.1-1.874 2.232-2.073C7.538 2 7.847 2 8.465 2c.27 0 .406 0 .536.011c.56.049 1.093.254 1.526.587c.1.078.196.167.388.344l.385.358c.571.53.857.795 1.198.972c.188.097.388.174.594.228c.377.1.78.1 1.588.1h.261c1.843 0 2.764 0 3.363.5c.055.046.108.095.157.146C19 5.802 19 6.658 19 8.369V9.8c0 2.451 0 3.677-.82 4.438c-.82.762-2.14.762-4.78.762h-2.8c-2.64 0-3.96 0-4.78-.761C5 13.476 5 12.25 5 9.8V5.217Z"/><path stroke-linecap="round" d="M22 20h-8M2 20h8m2-2v-3"/><circle cx="12" cy="20" r="2"/></g>`),
		g.Group(children),
	)
}