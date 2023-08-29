package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiBridge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6.95 20.95l-1.415-1.414a5 5 0 0 0 0-7.072L6.95 11.05a7 7 0 0 1 0 9.9Z"/><path fill="currentColor" d="m10.485 24.485l-1.414-1.414a10.001 10.001 0 0 0 0-14.142l1.414-1.414a12 12 0 0 1 0 16.97zm14.566-3.535a7 7 0 0 1 0-9.9l1.414 1.415a5 5 0 0 0 0 7.071z"/><path fill="currentColor" d="M21.515 24.485a12 12 0 0 1 0-16.97l1.414 1.414a10.001 10.001 0 0 0 0 14.142zM3 15H2V4H0v24h2V17h1a1 1 0 0 0 0-2zM30 4v11h-1a1 1 0 0 0 0 2h1v11h2V4z"/>`),
		g.Group(children),
	)
}