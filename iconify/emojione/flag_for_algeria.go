package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForAlgeria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M62 32C62 15.4 48.6 2 32 2v60c16.6 0 30-13.4 30-30"/><path fill="#699635" d="M2 32c0 16.6 13.4 30 30 30V2C15.4 2 2 15.4 2 32z"/><g fill="#ed4c5c"><path d="m40.7 39l-.1-5.4L46 32l-5.4-1.6l.1-5.4l-3.3 4.3l-5.4-1.6l3.3 4.3l-3.3 4.3l5.4-1.6z"/><path d="M37.2 44c-6.6 0-11.9-5.4-11.9-12s5.3-12 11.9-12c2.5 0 4.8.8 6.8 2.1C41.3 19 37.3 17 32.8 17C24.6 17 18 23.7 18 32s6.6 15 14.8 15c4.5 0 8.5-2 11.2-5.1c-1.9 1.3-4.2 2.1-6.8 2.1"/></g>`),
		g.Group(children),
	)
}