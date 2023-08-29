package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 34.512L256 477.488L0 34.512h41.601L256 405.502l214.399-370.99H512zM256 135.247L196.818 34.512h-89.766L256 292.248L404.948 34.512h-89.766L256 135.247z"/>`),
		g.Group(children),
	)
}