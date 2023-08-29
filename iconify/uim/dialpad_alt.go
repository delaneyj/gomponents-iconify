package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="4" height="4" x="10" y="6.955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="10" y=".955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="10" y="13.045" fill="currentColor" rx=".545"/><rect width="4" height="4" x="4" y="6.955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="4" y=".955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="4" y="13.045" fill="currentColor" rx=".545"/><rect width="4" height="4" x="16" y="6.955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="16" y=".955" fill="currentColor" rx=".545"/><rect width="4" height="4" x="16" y="13.045" fill="currentColor" rx=".545"/><rect width="4" height="4" x="10" y="19" fill="currentColor" rx=".545"/>`),
		g.Group(children),
	)
}