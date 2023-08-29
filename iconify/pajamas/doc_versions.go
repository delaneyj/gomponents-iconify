package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocVersions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 7v5.5h-7v-9H11V6a1 1 0 0 0 1 1h2.5Zm-2-3.379L14.379 5.5H12.5V3.621ZM7 2a1 1 0 0 0-1 1H4a1 1 0 0 0-1 1H1a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2a1 1 0 0 0 1 1h2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V5.414a1 1 0 0 0-.293-.707l-2.414-2.414A1 1 0 0 0 12.586 2H7Zm-1 9.5v-7H4.5v7H6Zm-3-1v-5H1.5v5H3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}