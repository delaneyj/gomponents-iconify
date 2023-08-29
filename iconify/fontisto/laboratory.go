package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laboratory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.171 19.68L14.819 8.369V2.962h1.708V0H6.098v2.965H7.82v5.407L.454 19.68A2.792 2.792 0 0 0 2.791 24h17.034a2.8 2.8 0 0 0 2.34-4.331l.007.011zm-.905 2.302a1.633 1.633 0 0 1-1.434.854H2.791a1.635 1.635 0 0 1-1.37-2.531l-.004.006l7.549-11.6V2.96h4.686v5.754l7.541 11.6c.17.251.272.561.272.895c0 .285-.074.553-.204.785l.004-.008z"/><path fill="currentColor" d="M14.412 12.351H8.221l-5.655 8.698a.287.287 0 0 0-.012.299l-.001-.001c.05.087.142.145.248.146h17.032a.283.283 0 0 0 .247-.145l.001-.001a.283.283 0 0 0-.013-.298l.001.001z"/>`),
		g.Group(children),
	)
}