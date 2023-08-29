package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Novelworld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.875 39.922l.1-28.09M10.056 8.078l-.05 25.434M7.553 11.134l.25 24.533m30.291-1.654l.05-25.935m-14.169 3.755c4.747-1.544 9.545-3.696 14.17-3.755m-28.089 0c5.512.488 9.355 2.437 13.919 3.755m-.047 25.583c4.555-2.305 9.349-2.938 14.166-3.403m-28.088-.501c5.588.449 9.6 2.316 13.922 3.904m16.619-26.082l-.2 24.633m-16.422 2.954a48.979 48.979 0 0 1 16.422-2.954m-32.544-.3c6.199.012 11.24 1.53 16.122 3.254m14.212-27.538l2.41-.05m-32.994-.199l2.496.351"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.842 9.807l.291 8.634l-1.552-1.602l-1.502 1.803l-.198-7.806"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.53 13.38l1.97-.043v25.384c-6.512-.58-12.298.72-18.625 1.202A84.636 84.636 0 0 0 5.5 38.57V13.537l2.077-.003"/>`),
		g.Group(children),
	)
}