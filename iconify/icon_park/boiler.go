package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boiler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="38" height="8" x="5" y="6" fill="#2F88FF" rx="2"/><path d="M8 14V40C8 41.1046 8.89543 42 10 42H38C39.1046 42 40 41.1046 40 40V14"/><path fill="#2F88FF" stroke-linecap="round" d="M31 29.0667C31 32.8958 27.866 36 24 36C20.134 36 17 32.8958 17 29.0667C17 25.2375 19.6923 23.7333 20.7692 20C25.0769 21.8666 25.0769 27.4667 25.0769 27.4667C25.0769 27.4667 26.1538 24.2667 29.3846 23.4667C29.6538 26.4 31 27.4316 31 29.0667Z"/></g>`),
		g.Group(children),
	)
}