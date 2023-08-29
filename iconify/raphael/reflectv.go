package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reflectv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.643 16.008v-.854h-1.705v.854h1.705zm3.41 0v-.854h-1.705v.854h1.705zm3.41 0v-.854h-1.705v.854h1.705zm2.596 0v-.854h-.892v.854h.89zm-12.828 0v-.854h-1.71v.854h1.71zm-13.64 0v-.854H1.89v.854h1.705zm3.41 0v-.854H5.3v.854h1.705zm3.412 0v-.854H8.71v.854h1.704zm3.41 0v-.854H12.12v.854h1.704zm-10.13-2.84h22.134V1.15L3.694 13.167zm3.354-.854l17.93-9.73v9.73H7.047zm18.78 5.728H3.694l22.134 12.015V18.042z"/>`),
		g.Group(children),
	)
}