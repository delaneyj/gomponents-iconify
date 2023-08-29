package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePalette0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePalette1" fill="currentColor"><path id="fePalette2" d="M10.525 21.892a4 4 0 0 0-6.77-4.232A9.954 9.954 0 0 1 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10c-.501 0-.994-.037-1.475-.108ZM9 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm6 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm3 5a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM6 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm9 5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`),
		g.Group(children),
	)
}