package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.435 0L0 6.052v12l10.435 5.953l10.435-6.052V5.948zm4.803 17.109l-.939.521l-.939-.521V12.94H7.515v4.174l-.94.521l-.94-.521V6.887l.939-.521l.939.521v4.174h5.843V6.887l.94-.521l.938.521z"/>`),
		g.Group(children),
	)
}