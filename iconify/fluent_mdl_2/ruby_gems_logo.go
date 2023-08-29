package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubyGemsLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M666 672h721l253 254q-154 155-306 307t-308 305L413 926l253-254zM1024 0l893 512v1024l-893 512l-893-512V512L1024 0zm723 1438V608l-723-417l-723 417v830l723 417l723-417z"/>`),
		g.Group(children),
	)
}