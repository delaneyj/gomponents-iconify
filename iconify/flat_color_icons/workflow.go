package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00BCD4" d="M7 31h10v10H7zm28.3-11.7l-5.6-5.6c-.4-.4-.4-1 0-1.4l5.6-5.6c.4-.4 1-.4 1.4 0l5.6 5.6c.4.4.4 1 0 1.4l-5.6 5.6c-.4.4-1 .4-1.4 0z"/><circle cx="12" cy="13" r="6" fill="#3F51B5"/><circle cx="36" cy="36" r="6" fill="#448AFF"/><g fill="#90A4AE"><path d="M11 24h2v5h-2z"/><path d="m12 21l-3 4h6z"/></g><g fill="#90A4AE"><path d="M20 12h5v2h-5z"/><path d="m28 13l-4-3v6z"/></g><g fill="#90A4AE"><path d="M35 21h2v5h-2z"/><path d="m36 29l3-4h-6z"/></g>`),
		g.Group(children),
	)
}