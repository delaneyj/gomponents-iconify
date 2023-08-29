package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhatsappSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5 4.768a.5.5 0 0 1 .761.34l.14.836a.5.5 0 0 1-.216.499l-.501.334A4.501 4.501 0 0 1 5 5.5v-.732ZM9.5 10a4.5 4.5 0 0 1-1.277-.184l.334-.5a.5.5 0 0 1 .499-.217l.836.14a.5.5 0 0 1 .34.761H9.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 3.253 6.182l-2.53 1.265a.5.5 0 0 1-.67-.67l1.265-2.53A7.467 7.467 0 0 1 0 7.5Zm4.23-3.42l.206-.138a1.5 1.5 0 0 1 2.311 1.001l.14.837a1.5 1.5 0 0 1-.648 1.495l-.658.438A4.522 4.522 0 0 0 7.286 9.42l.44-.658a1.5 1.5 0 0 1 1.494-.648l.837.14a1.5 1.5 0 0 1 1.001 2.311l-.138.207a.5.5 0 0 1-.42.229h-1A5.5 5.5 0 0 1 4 5.5v-1a.5.5 0 0 1 .23-.42Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}