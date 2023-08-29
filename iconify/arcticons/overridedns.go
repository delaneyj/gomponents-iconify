package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overridedns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.914 14.296L23.947 4.5L7.086 14.235l16.967 9.796l16.861-9.735z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.053 24.031L7.086 14.235v19.469L24.053 43.5V24.031zm16.861-9.735l-16.861 9.735V43.5l16.861-9.735V14.296zm-25.03 1.324l9.59-5.536l2.172 1.254c2.331 1.346 2.343 3.522.026 4.86l-1.199.692a9.303 9.303 0 0 1-8.417-.016Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.373 32.589V21.516l6.393 14.764V25.207m10.712 9.719c.592.542 1.334.443 2.366-.153l1.429-.825a5.31 5.31 0 0 0 2.407-4.158v0c0-1.529-1.078-2.146-2.407-1.378l-1.579.91c-1.33.769-2.407.151-2.407-1.377h0a5.31 5.31 0 0 1 2.407-4.159l1.429-.825c1.032-.596 1.774-.695 2.366-.152"/>`),
		g.Group(children),
	)
}