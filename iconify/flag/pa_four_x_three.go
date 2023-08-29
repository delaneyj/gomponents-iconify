package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPa4x30"><path fill-opacity=".7" d="M0 0h640v480H0z"/></clipPath></defs><g clip-path="url(#flagPa4x30)"><path fill="#fff" d="M0 0h640v480H0z"/><path fill="#fff" fill-rule="evenodd" d="M92.5 0h477.2v480H92.4z"/><path fill="#db0000" fill-rule="evenodd" d="M323 3.6h358v221.7H323z"/><path fill="#0000ab" fill-rule="evenodd" d="M3.2 225.3h319.9V480H3.2zm211.6-47.6l-42-29.4l-41.7 29.6l15.5-48L105 100l51.6-.4l16-48l16.3 47.9h51.6l-41.5 30l15.9 48z"/><path fill="#d80000" fill-rule="evenodd" d="m516.9 413.9l-42.4-27.7l-42.1 28l15.6-45.6l-42-28l52-.5l16.2-45.4l16.4 45.3h52l-41.8 28.5l16 45.4z"/></g>`),
		g.Group(children),
	)
}