package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M90.315 12.993H9.684a5.644 5.644 0 0 0-5.644 5.645V67.83a5.644 5.644 0 0 0 5.644 5.645h30.359v8.556h-9.402c-.892 0-1.613.721-1.613 1.612v1.751c0 .892.721 1.613 1.613 1.613h37.901c.891 0 1.613-.721 1.613-1.613v-1.751c0-.892-.722-1.612-1.613-1.612h-8.586v-8.556h30.359a5.643 5.643 0 0 0 5.645-5.645V18.638a5.645 5.645 0 0 0-5.645-5.645zM14.091 63.508V22.949h71.818v40.559H14.091z"/>`),
		g.Group(children),
	)
}