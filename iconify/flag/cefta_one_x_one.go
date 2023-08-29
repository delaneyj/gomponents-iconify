package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeftaOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#039" d="M0 0h512v512H0z"/><circle cx="256" cy="266.5" r="32.5" fill="none" stroke="#fc0" stroke-width="29.3"/><circle cx="256" cy="266.5" r="94.2" fill="none" stroke="#fc0" stroke-width="29.3"/><path fill="#039" d="m346.3 176.1l90.3 90.3l-90.3 90.3l-90.3-90.3z"/><path fill="#fc0" d="M102.1 251.8h63.2v29.3h-63.2zm276.4 0h94.2v29.3h-94.2zm-76.6-51.9l41.3-41.3l20.7 20.7l-41.3 41.3zM241.3 51.8h29.3V166h-29.3z"/><circle cx="154.8" cy="170.3" r="14.7" fill="#fc0"/><circle cx="68.6" cy="266.5" r="14.7" fill="#fc0"/><circle cx="256" cy="406.8" r="14.7" fill="#fc0"/><circle cx="256" cy="453.9" r="14.7" fill="#fc0"/><circle cx="350.2" cy="266.5" r="14.7" fill="#fc0"/><path fill="#fc0" d="m136.9 364.3l20.7-20.7l20.7 20.7l-20.7 20.7zm218.5 22.3L376 366l20.7 20.7l-20.7 20.8z"/>`),
		g.Group(children),
	)
}