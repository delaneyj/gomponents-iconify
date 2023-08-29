package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnalyzeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M4.99 12.862a7.1 7.1 0 0 0 12.171 3.924a1.956 1.956 0 0 1-.156-.637L17 16l.005-.15a2 2 0 1 1 1.769 2.137a9.099 9.099 0 0 1-15.764-4.85a1 1 0 0 1 1.98-.275z"/><path fill="currentColor" d="M12 8a4 4 0 1 1-3.995 4.2L8 12l.005-.2A4 4 0 0 1 12 8z"/><path fill="currentColor" d="M13.142 3.09a9.1 9.1 0 0 1 7.848 7.772a1 1 0 0 1-1.98.276a7.1 7.1 0 0 0-6.125-6.064A7.096 7.096 0 0 0 6.837 7.21a2 2 0 1 1-3.831.939L3 8l.005-.15a2 2 0 0 1 2.216-1.838a9.094 9.094 0 0 1 7.921-2.922z"/></g>`),
		g.Group(children),
	)
}