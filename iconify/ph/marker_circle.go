package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkerCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201.54 54.46A104 104 0 0 0 54.46 201.54A104 104 0 0 0 201.54 54.46ZM88 206.37V176h80v30.37a88.38 88.38 0 0 1-80 0ZM104 144h48v16h-48Zm14-16l10-48l10 48Zm72.22 62.23c-2 2-4.08 3.87-6.22 5.64V176a16 16 0 0 0-16-16v-16a16 16 0 0 0-13.61-15.8l-10.73-51.46a16 16 0 0 0-31.32 0l-10.73 51.46A16 16 0 0 0 88 144v16a16 16 0 0 0-16 16v19.87a88.917 88.917 0 0 1-6.22-5.64a88 88 0 1 1 124.44 0Z"/>`),
		g.Group(children),
	)
}