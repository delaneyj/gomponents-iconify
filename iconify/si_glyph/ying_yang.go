package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YingYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><path d="M7.984.053C3.599.053.045 3.614.045 8.006s3.555 7.953 7.939 7.953s7.939-3.561 7.939-7.953S12.369.053 7.984.053zM7.49 2.045c.838 0 1.519.654 1.519 1.46c0 .806-.681 1.46-1.519 1.46c-.84 0-1.519-.654-1.519-1.46c0-.806.679-1.46 1.519-1.46zm.545 12.863c-3.051 0-5.693-3.99.066-6.924c5.256-2.676 2.803-7.24-.066-6.976c4.375 0 6.939 3.111 6.939 6.95c.001 3.839-3.107 6.95-6.939 6.95z"/><ellipse cx="8.493" cy="11.445" rx="1.493" ry="1.445"/></g>`),
		g.Group(children),
	)
}