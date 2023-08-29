package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 92a8 8 0 1 1 8 8a8 8 0 0 1-8-8Zm-48-8a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm144 44a4 4 0 0 1-4 4h-28v12a83.64 83.64 0 0 1-3.87 25.2l25.47 11.13A4 4 0 0 1 232 188a4.09 4.09 0 0 1-1.6-.33l-25-10.95a84 84 0 0 1-154.72 0l-25 10.95A4.09 4.09 0 0 1 24 188a4 4 0 0 1-1.6-7.67l25.47-11.13A83.64 83.64 0 0 1 44 144v-12H16a4 4 0 0 1 0-8h28v-12a83.64 83.64 0 0 1 3.87-25.2L22.4 75.67a4 4 0 0 1 3.2-7.34l25 11a84 84 0 0 1 154.72 0l25-11a4 4 0 1 1 3.2 7.34L208.13 86.8A83.64 83.64 0 0 1 212 112v12h28a4 4 0 0 1 4 4Zm-192-4h152v-12a76 76 0 0 0-152 0Zm72 95.89V132H52v12a76.09 76.09 0 0 0 72 75.89ZM204 132h-72v87.89A76.09 76.09 0 0 0 204 144Z"/>`),
		g.Group(children),
	)
}