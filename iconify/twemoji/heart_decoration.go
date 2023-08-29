package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartDecoration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#DD2E44" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#F5F8FA" d="M31.885 13.764a7.66 7.66 0 0 0-7.66-7.661A7.65 7.65 0 0 0 18 9.309a7.646 7.646 0 0 0-6.224-3.206a7.661 7.661 0 0 0-7.661 7.661c0 .6.076 1.18.206 1.74C5.385 22.113 12.733 29.085 18 31c5.266-1.915 12.614-8.887 13.678-15.496c.131-.56.207-1.14.207-1.74z"/>`),
		g.Group(children),
	)
}