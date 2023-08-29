package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagNr1x10"><path fill-opacity=".7" d="M135.6 0h496.1v496h-496z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagNr1x10)" transform="translate(-140) scale(1.0321)"><path fill="#002170" d="M0 0h992.1v496H0z"/><path fill="#ffb20d" d="M0 226.8h992.1v42.4H0z"/><path fill="#fff" d="m292.4 424.4l-31.9-32l-10.2 44l-11.7-43.7l-30.9 33l11.8-43.6l-43.2 13l32-31.8l-44-10.3l43.6-11.6l-33-31l43.6 11.8l-13-43.2l31.8 32l10.3-44l11.7 43.6l30.8-32.9l-11.7 43.6l43.2-13l-32 31.8l44 10.3L290 362l33 30.9l-43.7-11.7z"/></g>`),
		g.Group(children),
	)
}