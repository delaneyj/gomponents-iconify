package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M21.125 0H4.875A4.874 4.874 0 0 0 0 4.875v16.25A4.874 4.874 0 0 0 4.875 26h16.25A4.874 4.874 0 0 0 26 21.125V4.875A4.874 4.874 0 0 0 21.125 0zM17.66 13.857l-6.229 4.949a.984.984 0 0 1-.951-.04a.966.966 0 0 1-.48-.816V8.051c0-.334.19-.644.479-.817a.986.986 0 0 1 .952-.039l6.229 4.948c.336.297.537.494.537.857s-.236.575-.537.857z"/>`),
		g.Group(children),
	)
}