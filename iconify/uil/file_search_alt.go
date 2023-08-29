package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearchAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7a1 1 0 0 0 0-2Zm0-14.59L14.59 8H13a1 1 0 0 1-1-1ZM7 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm14.71 12.29l-1.17-1.16A3.44 3.44 0 0 0 20 15a3.49 3.49 0 0 0-6 2.49a3.46 3.46 0 0 0 5.13 3.05l1.16 1.17a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-3.17-1.75a1.54 1.54 0 0 1-2.11 0a1.5 1.5 0 0 1-.43-1.05a1.46 1.46 0 0 1 .44-1.06a1.48 1.48 0 0 1 1-.43A1.47 1.47 0 0 1 19 17.49a1.5 1.5 0 0 1-.46 1.05ZM13 12H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Zm-2 6a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}