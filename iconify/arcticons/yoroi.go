package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yoroi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.885 20.904L42.5 7.719h-5.652l-12.852 8.756l-12.844-8.867H5.5l18.385 13.296zm3.659 2.874l11.074-7.748v4.207l-7.859 5.541l-3.215-2zm7.645 4.985l3.318-2.437v3.993l-.659.444l-2.659-2zm-3.882 6.867l2.993-2.104L9.271 16.03v3.985L31.307 35.63zm-4.318 2.881L9.374 26.326v3.993l14.4 10.073l3.215-1.881z"/>`),
		g.Group(children),
	)
}