package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cryptocom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.075 7.437L11.88 17.614l23.604.016l-2.305-10.192H14.075Zm-2.081 13.378L5.5 25.535l8.866 15.028H17.9l4.186-3.82V34.85l-4.345-4.067v-6.396l-5.747-3.57Zm18.275 3.617v6.397l-4.355 4.052l.016 1.925l4.21 3.757h3.506L42.5 25.58l-6.516-4.766l-5.715 3.617Zm-1.208-4.343H18.94l1.704 4.3l-.526 4.796l7.765-.012l-.487-4.784l1.665-4.3Z"/>`),
		g.Group(children),
	)
}