package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PandasIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#130754" d="M71.478 24.072h44.504v92.41H71.478v-92.41Zm0 189.825h44.504v92.409H71.478v-92.41ZM0 100.057h44.505v307.175H0V100.057Zm141.496 190.084H186v92.409h-44.504v-92.41Zm0-189.973H186v92.41h-44.504v-92.41ZM211.496 0H256v307.174h-44.505V0Z"/><path fill="#FFCA00" d="M71.478 143.454h44.505v43.6H71.478z"/><path fill="#E70488" d="M141.496 219.55h44.505v43.6h-44.505z"/>`),
		g.Group(children),
	)
}