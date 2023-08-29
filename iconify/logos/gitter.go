package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosGitter0" x1="50%" x2="50%" y1="0%" y2="100%"><stop offset="0%" stop-color="#FB0766"/><stop offset="100%" stop-color="#C50948"/></linearGradient></defs><path fill="url(#logosGitter0)" d="M0 0h256v256H0V0Z"/><path fill="#FFF" d="M83.914 62.873h12.525v82.661H83.914V62.873Zm76.149 20.039h12.524v62.622h-12.524V82.912Zm-50.599 0h12.524v110.466h-12.524V82.912Zm25.049 0h12.525v110.466h-12.525V82.912Z"/>`),
		g.Group(children),
	)
}