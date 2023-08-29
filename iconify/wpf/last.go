package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Last(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M21.125 0H4.875A4.874 4.874 0 0 0 0 4.875v16.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125V4.875A4.874 4.874 0 0 0 21.125 0zM19 18c0 .551-.449 1-1 1h-2c-.551 0-1-.449-1-1v-3.618l-5.569 4.425a.984.984 0 0 1-.951-.04a.969.969 0 0 1-.48-.817V8.051c0-.334.19-.644.479-.817a.986.986 0 0 1 .952-.039L15 11.618V8c0-.551.449-1 1-1h2c.551 0 1 .449 1 1v10z"/>`),
		g.Group(children),
	)
}