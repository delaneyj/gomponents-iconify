package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.84 12c0 2.05.985 3.225-.04 5c-1.026 1.775-2.537 1.51-4.314 2.534C14.71 20.56 14.184 22 12.133 22c-2.051 0-2.576-1.441-4.353-2.466C6.004 18.51 4.492 18.775 3.466 17c-1.025-1.775-.04-2.95-.04-5s-.985-3.225.04-5C4.492 5.225 6.003 5.49 7.78 4.466C9.556 3.44 10.082 2 12.133 2s2.577 1.441 4.353 2.466c1.776 1.024 3.288.759 4.313 2.534c1.026 1.775.04 2.95.04 5z"/>`),
		g.Group(children),
	)
}