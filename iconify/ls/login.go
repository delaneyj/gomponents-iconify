package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Login(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M391 672h195c37 0 68-13 92-38c24-24 39-56 39-92V151c0-36-15-68-39-92s-56-38-92-38H391v82h195c28 0 49 21 49 48v391c0 27-21 49-49 49H391v81zM0 269v156c0 18 15 33 33 33h182v123c0 11 5 20 15 25c4 1 9 1 11 1c7 0 13-2 18-7l235-235c9-9 8-27 0-37L259 94c-8-8-18-9-29-6c-10 5-15 13-15 24v124H33c-18 0-33 15-33 33z"/>`),
		g.Group(children),
	)
}