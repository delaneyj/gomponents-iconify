package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lipstick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#ed4c5c"><path d="m20.8 6.4l20.1 20.1l-11.2 11.2L2.2 10.2z"/><path d="M10.6 5.8c-5.2.9-9 2.7-8.6 4.1c.5 1.4 5.1 1.8 10.4.9c5.2-1 9.1-2.8 8.6-4.2c-.5-1.4-5.2-1.7-10.4-.8"/></g><path fill="#ff8080" d="M10.8 5.8c-4.9.8-8.5 2.6-8 3.9c.5 1.3 4.8 1.6 9.7.8c4.9-.8 8.5-2.6 8-3.8c-.5-1.4-4.8-1.8-9.7-.9"/><path fill="#e8e8e8" d="m16.855 26.363l12.869-12.869l16.829 16.829l-12.87 12.87z"/><path fill="#555b61" d="M44.407 54.577L57.84 41.14l4.173 4.171L48.58 58.748z"/><path fill="#b2c1c0" d="m28.047 38.228l13.43-13.439l3.89 3.888l-13.43 13.439z"/><path fill="#3e4347" d="m29.229 40.798l14.849-14.85L61.33 43.203l-14.85 14.85z"/>`),
		g.Group(children),
	)
}