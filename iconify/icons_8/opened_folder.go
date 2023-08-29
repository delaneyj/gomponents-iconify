package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenedFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 3v24.813l.78.156l12 2.5l1.22.25V28h6V15.437l1.72-1.718l.28-.314V3H5zm9.125 2H25v7.563l-1.72 1.718l-.28.314V26h-4v-8.906l-.28-.313L17 15.063V5.72l-.75-.19L14.125 5zM7 5.28l8 2v8.627l.28.313L17 17.937V28.28L7 26.188V5.282z"/>`),
		g.Group(children),
	)
}