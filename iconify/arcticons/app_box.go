package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" class="c"><path d="m40.9 14.273l-.055 19.454L24.061 43.5V24.047L40.9 14.273zm-16.839 9.774V43.5L7.1 33.727l.055-19.454l16.906 9.774z"/><path d="m40.9 14.273l-16.839 9.774l-16.906-9.774L23.973 4.5L40.9 14.273z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.083 6.761v20.486H13.56l10.501 10.476l10.175-10.464l-6.153-.012V6.873"/>`),
		g.Group(children),
	)
}