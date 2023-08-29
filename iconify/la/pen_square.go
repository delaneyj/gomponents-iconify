package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm13.127 1.877a3 3 0 0 0-2.125.875l-8.266 8.266l-.03.312l-.47 3.315l-.218 1.343l1.343-.218l3.313-.47l.312-.03l8.268-8.266c1.16-1.16 1.16-3.09 0-4.25a3.005 3.005 0 0 0-2.127-.877zm0 2c.254 0 .52.082.719.281a.978.978 0 0 1 .031 1.375l-.031.031l-7.797 7.829l-1.656.218l.218-1.656l7.828-7.797a.937.937 0 0 1 .688-.281z"/>`),
		g.Group(children),
	)
}