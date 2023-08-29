package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M86.304 40.985H13.696c-.836 0-1.513.677-1.513 1.513v15.127c0 .836.677 1.513 1.513 1.513h72.608c.836 0 1.513-.677 1.513-1.513V42.497c0-.835-.677-1.512-1.513-1.512zM56.261 17.848v-.053H13.696c-.836 0-1.513.677-1.513 1.513v15.127c0 .836.677 1.513 1.513 1.513h42.565v-.053a1.509 1.509 0 0 0 1.135-1.459V19.307c0-.704-.483-1.29-1.135-1.459zm9.622 46.205H13.696c-.836 0-1.513.677-1.513 1.513v15.127c0 .836.677 1.513 1.513 1.513h52.187c.836 0 1.513-.677 1.513-1.513V65.566c0-.836-.677-1.513-1.513-1.513z"/>`),
		g.Group(children),
	)
}