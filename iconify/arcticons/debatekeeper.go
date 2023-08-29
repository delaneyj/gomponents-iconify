package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Debatekeeper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.242 36.02a2.67 2.67 0 0 0-.592 1.691a2.709 2.709 0 0 0 2.708 2.709"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.358 7.58a2.707 2.707 0 0 0-2.251 4.212a9.889 9.889 0 0 0-7.64 9.625v9.304l-2.548 2.065a1.82 1.82 0 0 0 1.146 3.234h11.293m-8.36-26.599A15.816 15.816 0 0 0 4.5 19.886M19.358 7.58v32.84m0-14.292a7.831 7.831 0 0 0 0-15.663m0 18.666c7.34 0 12.423 1.628 15.247 5.43a2.567 2.567 0 0 1 .475 1.533v2.646a1.68 1.68 0 0 1-1.68 1.68H19.358m14.271-16.963a9.902 9.902 0 0 0 0-10.32m6.288 15.36a16.313 16.313 0 0 0 0-20.4"/>`),
		g.Group(children),
	)
}