package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27.812 16l-3.062-3.062V5.625h-2.625v4.688L16 4.188L4.188 16L7 15.933v11.942h17.875V16h2.937zM16 26.167h-5.833v-7H16v7zm5.667-3h-3.833v-4.042h3.833v4.042z"/>`),
		g.Group(children),
	)
}