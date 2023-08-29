package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minitonguewink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M951 731q0 122-75 207.5T695 1024t-181-85.5T439 731H73q-30 0-51.5-21T0 658.5t21.5-52T73 585h878q30 0 51.5 21.5t21.5 52t-21.5 51.5t-51.5 21zm-366 0q0 61 32 104t77.5 43t78-43T805 731H585zm293-438H658q-30 0-51.5-21.5t-21.5-52t21.5-52T658 146h220q30 0 51.5 21.5t21.5 52t-21.5 52T878 293zm-658.5 73Q159 366 116 312.5T73 183t43-129.5T219.5 0T323 53.5T366 183t-43 129.5T219.5 366z"/>`),
		g.Group(children),
	)
}