package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Meter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 16a9.928 9.928 0 0 0-1.14-4.618l-1.495 1.496A7.948 7.948 0 0 1 24 16zm-2.586-6L22 8.586L17.285 13.3A2.966 2.966 0 0 0 16 13a3 3 0 1 0 3 3a2.966 2.966 0 0 0-.3-1.285zM16 17a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1zm0-9a7.952 7.952 0 0 1 3.122.635l1.496-1.496A9.986 9.986 0 0 0 6 16h2a8.01 8.01 0 0 1 8-8z"/><path fill="currentColor" d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14Zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4Z"/>`),
		g.Group(children),
	)
}