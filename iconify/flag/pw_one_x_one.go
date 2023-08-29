package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PwOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPw1x10"><path fill-opacity=".7" d="M61.7 4.2h170.8V175H61.7z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagPw1x10)" transform="translate(-185 -12.5) scale(2.9973)"><path fill="#4aadd6" d="M0 4.2h301.2V175H0z"/><path fill="#ffde00" d="M185.9 86.8a52 52 0 0 1-53 50.8a52 52 0 0 1-53.2-50.8c0-28 23.8-50.8 53.1-50.8s53 22.7 53 50.8z"/></g>`),
		g.Group(children),
	)
}