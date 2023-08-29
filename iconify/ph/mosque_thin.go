package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MosqueThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 132a19.91 19.91 0 0 0-12 4v-8c0-39.63-28.82-58.82-52-74.23c-15-10-28-18.66-28-29.77a4 4 0 0 0-8 0c0 11.11-13 19.76-28 29.77C72.82 69.18 44 88.37 44 128v8a20 20 0 0 0-32 16v56a4 4 0 0 0 4 4h64a4 4 0 0 0 4-4v-32a12 12 0 0 1 24 0v32a4 4 0 0 0 4 4h32a4 4 0 0 0 4-4v-32a12 12 0 0 1 24 0v32a4 4 0 0 0 4 4h64a4 4 0 0 0 4-4v-56a20 20 0 0 0-20-20ZM100.41 60.43C112 52.71 122.51 45.71 128 37.14c5.49 8.57 16 15.57 27.59 23.29C177.41 75 202 91.34 203.87 124H52.13C54 91.34 78.59 75 100.41 60.43ZM20 152a12 12 0 0 1 24 0v52H20Zm140 4a20 20 0 0 0-20 20v28h-24v-28a20 20 0 0 0-40 0v28H52v-72h152v72h-24v-28a20 20 0 0 0-20-20Zm76 48h-24v-52a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}