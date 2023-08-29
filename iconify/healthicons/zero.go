package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 14a4 4 0 0 0-4 4v11.38l9.25-15.18A4 4 0 0 0 26 14h-4Zm7.96 3.436l-9.814 16.11A3.98 3.98 0 0 0 22 34h4a4 4 0 0 0 4-4V18c0-.191-.013-.38-.04-.564ZM14 18a8 8 0 0 1 8-8h4a7.98 7.98 0 0 1 5.334 2.037A7.985 7.985 0 0 1 34 18v12a8 8 0 0 1-8 8h-4c-2.37 0-4.5-1.033-5.962-2.666A7.977 7.977 0 0 1 14 30V18Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}