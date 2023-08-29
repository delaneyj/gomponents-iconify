package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningThreeBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm6 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm8 2a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/><path d="M12.75 10.145a1.994 1.994 0 0 0-1.5 0V5a.75.75 0 0 1 1.5 0v5.145Zm-1.5 3.71a1.994 1.994 0 0 0 1.5 0V19a.75.75 0 0 1-1.5 0v-5.145Zm8-1.71a1.994 1.994 0 0 1 1.5 0V5a.75.75 0 0 0-1.5 0v7.145Zm0 3.71V19a.75.75 0 0 0 1.5 0v-3.145a1.994 1.994 0 0 1-1.5 0Zm-14.5-7.71a1.994 1.994 0 0 0-1.5 0V5a.75.75 0 0 1 1.5 0v3.145Zm0 3.71a1.993 1.993 0 0 1-1.5 0V19a.75.75 0 0 0 1.5 0v-7.145Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}