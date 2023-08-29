package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Azurepipelines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.8 22.2H6V24H0v-6h1.8v4.2zM12.623 9.898l-7.635 7.635l1.479 1.479l7.634-7.636l-1.478-1.478zm-8.258 5.281l-.391-.39a1.504 1.504 0 0 1-.187-1.894L8.391 6H.998A1 1 0 0 0 0 7v5.5l3.502 3.604l.863-.924zM24 1v8.93a2 2 0 0 1-.89 1.664l-12.283 8.2a1 1 0 0 1-1.262-.124L8.04 18.146l6.768-6.77l-2.185-2.186l-6.77 6.77l-1.525-1.525a1 1 0 0 1-.125-1.262l8.2-12.284A2 2 0 0 1 14.067 0H23a1 1 0 0 1 1 1zm-3 5a3 3 0 1 0-6 0a3 3 0 0 0 6 0zM10.272 20.462c-.4 0-.778-.156-1.06-.439l-.369-.368l-.843.842l3.5 3.5H17a1 1 0 0 0 1-1v-7.39l-6.896 4.603a1.494 1.494 0 0 1-.832.253z"/>`),
		g.Group(children),
	)
}