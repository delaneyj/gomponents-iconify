package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.75 4.336v7l-7-7v15.328l7-7v7L20.414 12L12.75 4.336ZM17.586 12l-2.836 2.836V9.164L17.586 12Zm-7 0L7.75 14.836V9.164L10.586 12Z"/>`),
		g.Group(children),
	)
}