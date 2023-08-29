package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M331 213q25 0 56 7.5t56.5 24T469 283v58H0v-58q0-22 25.5-38.5t56.5-24t57-7.5q50 0 96 22q46-22 96-22zm-86 96v-26q0-10-35-24t-71.5-14T67 259t-35 24v26h213zm192 0v-26q0-10-35-24t-71-14q-32 0-65 12q11 12 11 26v26h160zM139 192q-31 0-53-22t-22-53t22-52.5T139 43t52.5 21.5T213 117t-21.5 53t-52.5 22zm-.5-117q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5zM331 192q-31 0-53-22t-22-53t22-52.5T331 43t52.5 21.5T405 117t-21.5 53t-52.5 22zm-.5-117q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5z"/>`),
		g.Group(children),
	)
}