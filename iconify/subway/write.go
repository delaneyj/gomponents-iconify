package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Write(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 439.5h170.7V268.8H0v170.7zm42.7-128H128v85.3H42.7v-85.3zm170.6 106.7H512v-42.7H213.3v42.7zM0 226.2h170.7V55.5H0v170.7zm42.7-128H128v85.3H42.7V98.2zm170.6-21.4v42.7H512V76.8H213.3zm0 256H512v-42.7H213.3v42.7zm0-128H512v-42.7H213.3v42.7z"/>`),
		g.Group(children),
	)
}