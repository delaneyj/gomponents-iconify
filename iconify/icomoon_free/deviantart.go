package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deviantart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.953 2.909V0h-2.909l-.291.294L8.378 2.91l-.431.291h-4.9v3.994h2.694l.241.291l-2.934 5.606v2.909h2.909l.291-.294l1.375-2.616l.431-.291h4.9V8.806H10.26l-.241-.294z"/>`),
		g.Group(children),
	)
}