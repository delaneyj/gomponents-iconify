package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M352.062 496h-192V295.993L16.047 296l.037-38.688L256.318 17.364L496 257.7v38.278l-143.938.006Zm-160-32h128V263.984l137.007-.006L256.274 62.636L54.672 264l137.39-.008Z"/>`),
		g.Group(children),
	)
}