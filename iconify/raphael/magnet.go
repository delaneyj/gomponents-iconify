package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.812 19.5h5.002v-6.867c-.028-1.706-.61-3.807-2.172-5.84c-1.54-2.015-4.315-3.72-7.94-3.688c-3.626-.032-6.402 1.674-7.94 3.687c-1.562 2.034-2.145 4.136-2.174 5.842V19.5h5v-6.866c-.027-.377.303-1.79 1.1-2.748c.818-.98 1.847-1.747 4.013-1.778c2.166.032 3.196.8 4.014 1.778c.798.96 1.126 2.372 1.1 2.748V19.5h-.002zm5.002 6.08V20.5h-5.002v5.08h5.004h-.002zm-20.226 0h5V20.5h-5v5.08z"/>`),
		g.Group(children),
	)
}