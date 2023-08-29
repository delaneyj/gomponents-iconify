package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Endpoints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.02 10.98h2v2h-2zm3.97 4.03h2v2h-2zm0-4.03h2v2h-2zm4.01 0h2v2h-2zm4.01 0h2v2h-2zm4.01 0h2v2h-2zm-4.01-3.97h2v2h-2zM5.99 17.99h4v4h-4zm8.02-15.98h4v4h-4z"/>`),
		g.Group(children),
	)
}