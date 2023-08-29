package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopStopwatchCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 8.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11ZM5.5 14a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Z"/><path d="M5.793 6.793a1 1 0 0 1 1.414 0l1.5 1.5a1 1 0 0 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414Zm11.5 2.914a1 1 0 0 0 1.414 0l1.5-1.5a1 1 0 0 0-1.414-1.414l-1.5 1.5a1 1 0 0 0 0 1.414Zm-1.013 1.668a1 1 0 0 1-.155 1.406l-2.5 2a1 1 0 0 1-1.25-1.562l2.5-2a1 1 0 0 1 1.406.156ZM12 7V5h2v2h-2Z"/><path d="M10.5 5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopStopwatchCircleFilled0)"/></g>`),
		g.Group(children),
	)
}