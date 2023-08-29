package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOnLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M9 22v-4.343c0-.818 0-1.226-.152-1.594c-.152-.367-.442-.656-1.02-1.235l-3.242-3.242c-.29-.29-.434-.434-.51-.617C4 10.785 4 10.58 4 10.171V10c0-.943 0-1.414.293-1.707C4.586 8 5.057 8 6 8h12c.943 0 1.414 0 1.707.293C20 8.586 20 9.057 20 10v.172c0 .408 0 .613-.076.796c-.076.184-.22.329-.51.618l-3.242 3.242c-.579.579-.867.867-1.02 1.235c-.152.368-.152.776-.152 1.594V22m0-6H9"/><path d="M4.5 11h15M12 5V2" opacity=".5"/><path d="M8 5L6 3m10 2l2-2"/><path d="M12 19v2" opacity=".5"/></g>`),
		g.Group(children),
	)
}