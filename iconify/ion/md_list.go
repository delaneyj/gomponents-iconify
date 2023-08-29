package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M80 280h256v48H80z" fill="currentColor"/><path d="M80 184h320v48H80z" fill="currentColor"/><path d="M80 88h352v48H80z" fill="currentColor"/><g><path d="M80 376h288v48H80z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}