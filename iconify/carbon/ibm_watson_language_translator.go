package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmWatsonLanguageTranslator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 28h-3c-3.9 0-7-3.1-7-7v-4h2v4c0 2.8 2.2 5 5 5h3v2zm12 2h2.2l-4.6-11h-2.2l-4.6 11H21l.8-2h5.3l.9 2zm-5.3-4l1.8-4.4l1.8 4.4h-3.6zM28 15h-2v-4c0-2.8-2.2-5-5-5h-4V4h4c3.9 0 7 3.1 7 7v4zM14 5V3H9V1H7v2H2v2h8.2c-.2.9-.8 2.5-2.2 4c-.6-.7-1.1-1.4-1.4-2H4.3c.4 1 1.1 2.2 2.1 3.3c-.8.7-2 1.3-3.4 1.8l.7 1.9c1.8-.7 3.2-1.5 4.3-2.3c1.1.9 2.5 1.7 4.3 2.3l.7-1.9c-1.4-.5-2.6-1.2-3.5-1.8c1.9-2 2.5-4.1 2.7-5.3H14z"/>`),
		g.Group(children),
	)
}