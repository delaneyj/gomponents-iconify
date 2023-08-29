package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5.105 12c-.397 0-.681-.088-.853-.264c-.168-.18-.252-.465-.252-.853V4.117c0-.397.086-.681.258-.853c.176-.176.458-.264.847-.264H9.03c1.108 0 2.025.982 2.025 2.185c0 .9-.45 1.634-1.35 2.051c1.162.213 1.814 1.392 1.814 2.245c0 1.031-.528 2.519-2.24 2.519H5.104Zm3.274-3.997H5.8v2.628h2.579c.521 0 1.25-.51 1.25-1.332c0-.823-.729-1.296-1.25-1.296ZM5.8 4.37v2.327h2.38c.36 0 1.097-.337 1.097-1.196c0-.86-.797-1.131-1.097-1.131H5.8Z"/>`),
		g.Group(children),
	)
}