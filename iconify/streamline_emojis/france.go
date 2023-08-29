package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func France(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ff6242" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M45 12.25h-6.32A45.89 45.89 0 0 1 24 9.88A45.73 45.73 0 0 0 9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37A46 46 0 0 1 24 37.46a45.62 45.62 0 0 0 14.65 2.38H45a.93.93 0 0 0 1-.79V13a.94.94 0 0 0-1-.75Z"/><path fill="#fff" stroke="#45413c" d="M24 9.88A43.57 43.57 0 0 0 16.28 8v27.61A44.44 44.44 0 0 1 24 37.46a43.26 43.26 0 0 0 7.74 1.86V11.73A44.11 44.11 0 0 1 24 9.88Z"/><path fill="#009fd9" stroke="#45413c" d="M9.37 7.5H3c-.58 0-1 .35-1 .79v26a.94.94 0 0 0 1 .79h6.37a46.66 46.66 0 0 1 6.91.52V8a46.66 46.66 0 0 0-6.91-.5Z"/><path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/>`),
		g.Group(children),
	)
}