package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuUd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<mask id="circleFlagsRuUd0"><circle cx="256" cy="256" r="256" fill="#fff"/></mask><g mask="url(#circleFlagsRuUd0)"><path fill="#eee" d="M160 0h192l32 256l-32 256H160l-32-256Z"/><path fill="#333" d="M0 0h160v512H0Z"/><path fill="#d80027" d="M352 0h160v512H352ZM229 176l107 107H176l107-107v160L176 229h160L229 336Z"/></g>`),
		g.Group(children),
	)
}