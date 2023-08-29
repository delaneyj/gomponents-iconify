package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandItch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7v1c0 1.087 1.078 2 2 2c1.107 0 2-.91 2-2c0 1.09.893 2 2 2s2-.91 2-2c0 1.09.893 2 2 2s2-.91 2-2c0 1.09.893 2 2 2s2-.91 2-2c0 1.09.893 2 2 2c.922 0 2-.913 2-2V7c-.009-.275-.538-.964-1.588-2.068A3 3 0 0 0 18.238 4H5.762a3 3 0 0 0-2.174.932C2.538 6.036 2.008 6.725 2 7zm2 3c-.117 6.28.154 9.765.814 10.456c1.534.367 4.355.535 7.186.536c2.83-.001 5.652-.169 7.186-.536c.99-1.037.898-9.559.814-10.456"/><path d="m10 16l2-2l2 2m-2-2v4"/></g>`),
		g.Group(children),
	)
}