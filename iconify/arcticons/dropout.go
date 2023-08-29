package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.1 40.906V7.094h5.494C34.892 7.094 42.5 14.702 42.5 24h0c0 9.298-7.608 16.906-16.906 16.906H20.1ZM5.5 26.95h9.241v9.241H5.5zm0-15.967h9.241v9.241H5.5z"/>`),
		g.Group(children),
	)
}