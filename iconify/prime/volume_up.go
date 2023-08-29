package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19.75a.81.81 0 0 1-.47-.16l-4.79-3.84H3a.76.76 0 0 1-.75-.75V9A.76.76 0 0 1 3 8.25h4.74l4.79-3.84a.75.75 0 0 1 1.22.59v14a.76.76 0 0 1-.43.68a.71.71 0 0 1-.32.07Zm-9.25-5.5H8a.78.78 0 0 1 .47.16l3.78 3V6.56l-3.78 3a.78.78 0 0 1-.47.19H3.75Zm14.71 3.82a.76.76 0 0 1-.49-.18a.74.74 0 0 1-.07-1.06a7.24 7.24 0 0 0 0-9.66a.75.75 0 1 1 1.12-1a8.7 8.7 0 0 1 0 11.64a.72.72 0 0 1-.56.26Z"/><path fill="currentColor" d="M16.11 15.38a.78.78 0 0 1-.45-.15a.75.75 0 0 1-.15-1.05a3.6 3.6 0 0 0 0-4.36a.75.75 0 0 1 1.2-.9a5.07 5.07 0 0 1 0 6.16a.77.77 0 0 1-.6.3Z"/>`),
		g.Group(children),
	)
}