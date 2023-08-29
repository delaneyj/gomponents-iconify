package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuddiconsForums(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 7h-7C5.67 7 5 6.33 5 5.5S5.67 4 6.5 4h1.59A1.64 1.64 0 0 1 8 3.5C8 2.67 8.67 2 9.5 2h1c.83 0 1.5.67 1.5 1.5c0 .18-.04.34-.09.5h1.59c.83 0 1.5.67 1.5 1.5S14.33 7 13.5 7zM4 8h12c.55 0 1 .45 1 1s-.45 1-1 1H4c-.55 0-1-.45-1-1s.45-1 1-1zm1 3h10c.55 0 1 .45 1 1s-.45 1-1 1H5c-.55 0-1-.45-1-1s.45-1 1-1zm2 3h6c.55 0 1 .45 1 1s-.45 1-1 1h-1.09c.05.16.09.32.09.5c0 .83-.67 1.5-1.5 1.5h-1c-.83 0-1.5-.67-1.5-1.5c0-.18.04-.34.09-.5H7c-.55 0-1-.45-1-1s.45-1 1-1z"/>`),
		g.Group(children),
	)
}