package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.531 4.822c-.747 3.316-3.053 5.066-6.688 5.066H6.634l-.841 5.338H4.78l-.053.344a.374.374 0 0 0 .369.431h2.588a.625.625 0 0 0 .616-.525l.025-.131l.488-3.091l.031-.169a.622.622 0 0 1 .616-.525h.384c2.506 0 4.469-1.019 5.044-3.963c.216-1.119.134-2.069-.356-2.775z"/><path fill="currentColor" d="M12.984 1.206C12.243.362 10.903 0 9.19 0H4.218c-.35 0-.65.253-.703.6L1.443 13.731a.427.427 0 0 0 .422.494h3.072l.772-4.891l-.025.153c.053-.347.35-.6.7-.6h1.459c2.866 0 5.109-1.162 5.766-4.531c.019-.1.037-.197.05-.291c.194-1.244 0-2.094-.675-2.859z"/>`),
		g.Group(children),
	)
}