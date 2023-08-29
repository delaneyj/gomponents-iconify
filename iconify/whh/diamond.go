package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m896 0l128 256H768zM576 0h256L704 256zM384 256L512 0l128 256H384zM192 0h256L320 256zM0 256L128 0l128 256H0zm448 640L0 320h256zm256-576L512 896L320 320h384zm320 0L576 896l192-576h256z"/>`),
		g.Group(children),
	)
}