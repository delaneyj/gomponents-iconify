package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellZFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M14.235 19c.865 0 1.322 1.024.745 1.668A3.992 3.992 0 0 1 12 22a3.992 3.992 0 0 1-2.98-1.332c-.552-.616-.158-1.579.634-1.661l.11-.006h4.471zM12 2c1.358 0 2.506.903 2.875 2.141l.046.171l.008.043a8.013 8.013 0 0 1 4.024 6.069l.028.287L19 11v2.931l.021.136a3 3 0 0 0 1.143 1.847l.167.117l.162.099c.86.487.56 1.766-.377 1.864L20 18H4c-1.028 0-1.387-1.364-.493-1.87a3 3 0 0 0 1.472-2.063L5 13.924l.001-2.97A8 8 0 0 1 8.822 4.5l.248-.146l.01-.043a3.003 3.003 0 0 1 2.562-2.29l.182-.017L12 2zm2 6h-4l-.117.007A1 1 0 0 0 9 9l.007.117A1 1 0 0 0 10 10h1.584l-2.291 2.293l-.076.084C8.703 13.014 9.147 14 10 14h4l.117-.007A1 1 0 0 0 15 13l-.007-.117A1 1 0 0 0 14 12h-1.586l2.293-2.293l.076-.084C15.297 8.986 14.853 8 14 8z"/></g>`),
		g.Group(children),
	)
}