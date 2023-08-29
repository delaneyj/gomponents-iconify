package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TVMonitorSelected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1716 1280l-128 128h-564v256H640v-128h256v-128H0V256h1920v820l-128 128V384H128v896h1588zm223 3l90 91l-557 557l-269-269l90-91l179 179l467-467z"/>`),
		g.Group(children),
	)
}