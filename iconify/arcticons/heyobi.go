package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heyobi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.803 30.688v5.698H10.87l-5.37 4.61V7.004h35.303v6.13m-1.529 14.162H42.5m-3.226-10.752H42.5m-1.613 0v10.752"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.344 23.667a3.562 3.562 0 1 0 7.123 0v-3.629a3.614 3.614 0 0 0-3.628-3.628a3.502 3.502 0 0 0-3.495 3.629Zm15.42-1.747a2.688 2.688 0 0 1 0 5.376h-4.435V16.544h4.435a2.688 2.688 0 0 1 0 5.376Zm0 0h-4.435"/>`),
		g.Group(children),
	)
}