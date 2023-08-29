package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.922 5.626A3.72 3.72 0 0 0 10.205 2a3.712 3.712 0 0 0-2.92 1.418a2.09 2.09 0 0 0-3.719 1.573a3.028 3.028 0 0 0-3.567 2.98A3.028 3.028 0 0 0 3.026 11H4.46l3.539 3.664L11.538 11h1.742a2.725 2.725 0 0 0 .641-5.374zM8 13l-3-3h2V7h2v3h2l-3 3z"/>`),
		g.Group(children),
	)
}