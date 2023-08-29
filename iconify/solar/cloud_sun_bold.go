package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSunBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.286 20C19.442 20 22 17.472 22 14.353c0-2.472-1.607-4.573-3.845-5.338C17.837 6.194 15.415 4 12.476 4C9.32 4 6.762 6.528 6.762 9.647c0 .69.125 1.35.354 1.962a4.351 4.351 0 0 0-.83-.08C3.919 11.53 2 13.426 2 15.765C2 18.104 3.919 20 6.286 20h10Z"/><path d="M9.94 2.955a5 5 0 0 0-6.327 7.723a5.788 5.788 0 0 1 1.664-.561a7.169 7.169 0 0 1-.015-.47c0-3.073 1.951-5.677 4.678-6.692Z"/></g>`),
		g.Group(children),
	)
}