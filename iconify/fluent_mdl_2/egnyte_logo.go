package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EgnyteLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1843 662l-663 488v770H870v-768q-1-1-28-21t-71-52t-100-75t-117-86t-119-87t-107-78t-81-60t-42-31l262-193l557 410l557-411l262 194zm-431 599l418 308l-92 125l-418-308l92-125zm-309-613H947V128h156v520zm-467 613l92 125l-418 308l-92-125l418-308z"/>`),
		g.Group(children),
	)
}