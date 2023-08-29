package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DestinyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="8.192" height="15.019" x="19.904" y="8.764" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.096"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.897 10.209a25.731 25.731 0 0 1-6.546-2.494c-.193-.136-.873-.495-.873-.495a5.622 5.622 0 0 0-5.55 9.73c5.428 4.096 10.44 11.384 10.44 18.723a5.632 5.632 0 0 0 11.264 0c0-7.339 5.012-14.627 10.44-18.722a5.62 5.62 0 0 0-5.55-9.731s-.68.359-.873.494a25.73 25.73 0 0 1-6.546 2.494"/>`),
		g.Group(children),
	)
}