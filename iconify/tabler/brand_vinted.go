package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVinted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.028 6c0 7.695-.292 11.728 0 12c2.046-5 4.246-12.642 5.252-14.099c.343-.497.768-.93 1.257-1.277c.603-.39 1.292-.76 1.463-.575c-.07 2.319-4.023 15.822-4.209 16.314a6.135 6.135 0 0 1-3.465 3.386c-3.213.78-3.429-.446-3.836-1.134C6.54 18.512 5.808 6.355 6.045 5c.05-.523.143-1.851 2.491-2c2.359-.354 2.547 1.404 2.492 3z"/>`),
		g.Group(children),
	)
}