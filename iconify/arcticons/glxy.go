package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glxy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.099 13.092a18.815 18.815 0 0 1-6.256 27.226m-5.428 2.004a18.882 18.882 0 0 1-12.637-1.875M6.026 26.202A18.823 18.823 0 0 1 24.692 5.074a18.743 18.743 0 0 1 7.587 1.593"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.51 9.467a4.854 4.854 0 0 1-4.199 4.81c-2.386.335-9.314 6.622-13.316 8.893c3.96-4.091 7.887-10.65 7.81-13.703a4.752 4.752 0 0 1 4.853-4.853a4.853 4.853 0 0 1 4.852 4.853ZM22.082 31.538c.228 4.475 14.287 10.579 19.71 11.836c-13.011.27-21.51-4.043-26.611-3.248a11.94 11.94 0 0 1-1.89.204a8.791 8.791 0 1 1 8.791-8.792Z"/>`),
		g.Group(children),
	)
}