package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LcOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<g fill-rule="evenodd"><path fill="#65cfff" d="M0 0h512v512H0z"/><path fill="#fff" d="m254.8 44.8l173.5 421.6l-344 1L254.7 44.8z"/><path d="m255 103l150 362.6l-297.5.8L255 103z"/><path fill="#ffce00" d="m254.8 256.1l173.5 210.8l-344 .5l170.5-211.3z"/></g>`),
		g.Group(children),
	)
}