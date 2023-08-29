package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cherries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.702 3.259C9.947 5.103 7.555 13.695 7.592 17.063A6.501 6.501 0 0 0 8.5 30a6.5 6.5 0 0 0 .963-12.93c.179-3.174 1.372-5.73 2.969-7.625l.005-.007c.493-.585.834-.99 1.333-.223l5.895 9.059a6.5 6.5 0 1 0 1.671-.894l-6.027-9.283c-.493-.765-.325-.964.277-1.36a11.718 11.718 0 0 1 2.419-1.134c.694-.224.862.128.935.302c1.76 4.226 3.826 5.872 5.004 6.356c1.83.692 3.193.927 5.623.1c.234-.08.316-.345.129-.647c0 0-1.64-2.567-3.236-5.421c0 0-1.822-3.807-6.758-3.034ZM8.5 28a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm19.53-4.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}