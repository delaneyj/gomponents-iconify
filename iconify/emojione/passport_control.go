package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassportControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><g fill="#fff"><path d="M34.3 23c0 4.8 9.6 4.8 9.6 0v-4.8h-9.6l-3.8 2.9h3.8V23m9.5-5.7v-3.8H33.3l1 3.8zm-19.2 8.3l-10.3-2.2v9.9l10.3 2.1zm0-6.5l-10.3 2.1l10.3 2.2z"/><path d="m30.9 43.1l14.4-14.4H30.9l-6.7 11.9l-1.9-3.4h-5.7l4.8 8.6h5.7l3.8-6.9zm1.7.7h15.1V28.7zm-1.7 1.7h16.8v5H30.9z"/></g>`),
		g.Group(children),
	)
}