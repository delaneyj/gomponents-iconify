package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PwFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPw4x30"><path fill-opacity=".7" d="M-70.3 0h640v480h-640z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagPw4x30)" transform="translate(70.3)"><path fill="#4aadd6" d="M-173.4 0h846.3v480h-846.3z"/><path fill="#ffde00" d="M335.6 232.1a135.9 130.1 0 1 1-271.7 0a135.9 130.1 0 1 1 271.7 0z"/></g>`),
		g.Group(children),
	)
}