package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BytedanceMiniApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path stroke="#fff" d="M26.1215 21.8789L21.8789 26.1215"/><path stroke="#fff" d="M32.4851 26.8285L33.8993 25.4143C35.4614 23.8522 35.4614 21.3195 33.8993 19.7574L28.2425 14.1005C26.6804 12.5384 24.1477 12.5384 22.5856 14.1005L21.1714 15.5148"/><path stroke="#fff" d="M26.8282 32.4853L25.414 33.8995C23.8519 35.4616 21.3193 35.4616 19.7572 33.8995L14.1003 28.2427C12.5382 26.6806 12.5382 24.1479 14.1003 22.5858L15.5145 21.1716"/></g>`),
		g.Group(children),
	)
}