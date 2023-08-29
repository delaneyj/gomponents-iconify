package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandZoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.45 3.333c-.801 0-1.45.65-1.45 1.45v4.986c0 1.6 1.298 2.898 2.898 2.898h6.986c.8 0 1.45-.649 1.45-1.45V6.233a2.9 2.9 0 0 0-2.899-2.899H1.45ZM16 4.643v6.715c0 .544-.618.86-1.059.539l-2.059-1.498a1.333 1.333 0 0 1-.549-1.078V6.679c0-.427.204-.827.55-1.078l2.058-1.498a.667.667 0 0 1 1.059.54Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}